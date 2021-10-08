package DataStruct

//定义匿名空结构体
var (
	defSize = 16
	emptyElement = struct {}{}
)


//使用map作为set的数据结构 interface作为key，空结构体作为value
//不可重复，无序
//不是线程安全的
type HashSet struct {
	set map[interface{}]struct{}
}

func NewHashSet(size int) *HashSet {
	if size < 16 {
		size = 16
	}
	return &HashSet{
		set: make(map[interface{}]struct{},size),
	}
}

func (hs *HashSet) Add(ele ...interface{})  {
	for i := range ele {
		hs.set[ele[i]] = emptyElement
	}
}

func (hs *HashSet) Remove(ele ...interface{})  {
	for i := range ele {
		delete(hs.set, ele[i])
	}
}

func (hs *HashSet) Contains(ele interface{}) bool {
	_,ok := hs.set[ele]
	return ok
}

func (hs *HashSet) Size() int {
	return len(hs.set)
}

func (hs *HashSet) IsEmpty()bool  {
	return len(hs.set) == 0
}

func (hs *HashSet) Clear() {
	hs.set = make(map[interface{}]struct{}, defSize)
}

func (hs *HashSet) Iterator() []interface{} {
	res := make([]interface{},0,len(hs.set))
	for k,_ := range hs.set {
		res = append(res,k)
	}
	return res
}
