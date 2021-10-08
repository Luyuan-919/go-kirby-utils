package DataStruct

type kmap struct {
	positive map[interface{}]interface{}
	reverse map[interface{}]interface{}
}

//创建双向map
func NewKmap() *kmap {
	return &kmap{
		positive: make(map[interface{}]interface{}),
		reverse: make(map[interface{}]interface{}),
	}
}

//存值
func (k *kmap) Set(key,value interface{}) {
	k.set(key,value)
}

func (k *kmap) set(key,value interface{}) {
	k.positive[key] = value
	k.reverse[value] = key
}

//根据key获取value
func (k *kmap) Get(key interface{}) interface{} {
	return k.getValue(key)
}

func (k *kmap) getValue(key interface{}) interface{} {
	return k.positive[key]
}

//根据value获取key
func (k *kmap) GetKey(value interface{}) interface{} {
	return k.getKey(value)
}

func (k *kmap) getKey(value interface{}) interface{} {
	return k.reverse[value]
}

//判断value是否存在
func (k *kmap) ValueIsExist(key interface{}) bool {
	_, ok := k.positive[key]
	return ok
}

//判断key是否存在
func (k *kmap) KeyIsExist(value interface{}) bool {
	_, ok := k.reverse[value]
	return ok
}