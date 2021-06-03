package kcache

import (
	"container/list"
	"sync"
)

const (
	MB = 1024 * 1024
	DefMaxMemory = 32 * MB
	SizeofStringMemory = 16
)


type cache struct {
	//最大内存
	maxMemory int
	//已经使用过的内存
	usedMemory int
	//缓存的map 值是string 而value是双向链表中对应节点的指针
	cache map[string]*list.Element
	//缓存的value指向的双向链表
	ll *list.List
	//当缓存中的某条数据要被清理时，OnEvicted作为回调函数 这个函数可以是nil
	OnEvicted func(key string,value Value)
}

//允许值是任意 实现了Size()方法的类型，这个函数是返回值所占用的类型的大小
type Value interface {
	Size() int
}

//键值对 entry 是双向链表节点的数据类型，在链表中仍保存每个值对应的 key 的好处在于，淘汰队首节点时，需要用 key 从字典中删除对应的映射
type entry struct {
	key string
	value Value
}

func newCache(onEvicted func(string,Value))*cache {
	return &cache{
		maxMemory: DefMaxMemory,
		ll: list.New(),
		cache: make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

//实例化Cache
func newCacheWithMaxMemory(maxMemory int,onEvicted func(string,Value))*cache {
	//当传入的最大内存数非法时，进行修正，使用默认的最大内存
	if maxMemory <= 0 || maxMemory > DefMaxMemory{
		maxMemory = DefMaxMemory
	}
	return &cache{
		maxMemory: maxMemory * MB,
		ll: list.New(),
		cache: make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

//查找
func (c *cache) get(key string) (value Value,ok bool) {
	//如果缓存中存在key
	if ele, ok := c.cache[key]; ok {
		//双向链表中把查到的元素移动到队尾 即 缓存查到 提高该元素的优先级
		c.ll.MoveToFront(ele)
		//断言查到的元素 通过查到的元素得到entry类型的值
		kv := ele.Value.(*entry)
		return kv.value,true
	}
	return
}

//缓存淘汰
func (c *cache) reMoveOldest() {
	//拿到双向链表中的队首元素
	ele := c.ll.Back()
	//如果ele不为nil  也就是说双向链表的长度不为0
	if ele != nil {
		//移除该元素
		c.ll.Remove(ele)
		//拿到该元素的*entry值
		kv := ele.Value.(*entry)
		//在缓存数组中删除该key对应的位置
		delete(c.cache,kv.key)
		//这里删除缓存后，需要更新缓存中的已经使用过的内存大小
		//而删除的队首元素的key 是一个string 其内存占用大小恒为16  直接写进const 进行调用
		//而value占用的内存大小 就需要在使用的时候，自己实现一下value接口
		//在实现接口的时候，只需要调用Memory包中的Memory.SizeofMemoryInt()方法，即可获得该类型的内存大小
		c.usedMemory -= SizeofStringMemory + kv.value.Size()
		//如果回调函数不为空，则执行回调函数
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key,kv.value)
		}
	}
}
func (c *cache) delete(key string){
	//如果元素不存在，就说明缓存里没有这个值，因此直接返回true即可
	if _, ok := c.cache[key]; !ok {
		return
	}
	//如果存在
	v,_ := c.cache[key]
	//移除该元素
	c.ll.Remove(v)
	//拿到该元素的*entry值
	kv := v.Value.(*entry)
	//在缓存数组中删除该key对应的位置
	delete(c.cache,kv.key)
	//这里删除缓存后，需要更新缓存中的已经使用过的内存大小
	//而删除的队首元素的key 是一个string 其内存占用大小恒为16  直接写进const 进行调用
	//而value占用的内存大小 就需要在使用的时候，自己实现一下value接口
	//在实现接口的时候，只需要调用Memory包中的Memory.SizeofMemoryInt()方法，即可获得该类型的内存大小
	c.usedMemory -= SizeofStringMemory + kv.value.Size()
	//如果回调函数不为空，则执行回调函数
	if c.OnEvicted != nil {
		c.OnEvicted(kv.key,kv.value)
	}
	return
}


//增加/更新
func (c *cache) add(key string, value Value) {
	//如果元素存在，则更新，并加到队尾
	if _, ok := c.cache[key]; ok {
		c.update(key,value)
	}else {
		//元素不存在，就新建节点加入链表
		ele := c.ll.PushFront(&entry{
			key: key,
			value: value,
		})
		c.cache[key] = ele
		c.usedMemory += SizeofStringMemory + value.Size()
	}
	//如果超出最大内存 则删除最不活跃的节点 一直到最大内存未超
	for c.maxMemory != 0 && c.usedMemory > c.maxMemory {
		c.reMoveOldest()
	}
}

func (c *cache) update(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		//现在的kv，还存的是原来的数据 先已经使用的内存减去现在元素的大小
		c.usedMemory -= kv.value.Size()
		//更新value
		kv.value = value
		//把现在占用的内存加上
		c.usedMemory += kv.value.Size()
	}else{
		c.add(key,value)
	}
}



//封装Cache  实现并发
type Kcache struct {
	sync.RWMutex
	c *cache
}

func NewKcache(onEvicted func(string,Value))*Kcache  {
	return &Kcache{
		c: newCache(onEvicted),
	}
}

func NewKcacheWithMaxMemory(MaxMemory int,onEvicted func(string,Value))*Kcache  {
	return &Kcache{
		c: newCacheWithMaxMemory(MaxMemory,onEvicted),
	}
}

func (k *Kcache) Add(key string, value Value) {
	k.Lock()
	defer k.Unlock()
	k.c.add(key,value)
}

func (k *Kcache) Get(key string) (v Value,ok bool) {
	k.RLock()
	defer k.RUnlock()
	if v, ok = k.c.get(key); ok {
		return
	}
	return
}

func (k *Kcache) Update(key string, value Value) {
	k.Lock()
	defer k.Unlock()
	k.c.update(key,value)
}

func (k *Kcache) Delete(key string)  {
	k.Lock()
	defer k.Unlock()
	k.c.delete(key)
}