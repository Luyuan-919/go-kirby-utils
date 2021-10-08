package DataStruct

const (
	DefCapOfStack = 64
)

type Stack struct {
	elements []interface{}
	len int
	Cap int
}

//新建一个栈
func NewStack() *Stack {
	return &Stack{
		elements: make([]interface{},64),
		len: 0,
		Cap: 64,
	}
}

func NewStackWithValues(ele []int) *Stack {
	s := NewStack()
	for _,v := range ele{
		s.Push(v)
	}
	return s
}

//带容量的栈
func NewStackWithCap(cap int) *Stack  {
	if cap <= 0 {
		cap = DefCapOfStack
	}
	return &Stack{
		elements: make([]interface{},cap),
		len: 0,
		Cap: cap,
	}
}

//当栈的长度和容量相同时，扩容
//长度小于1024的时候 1.5倍扩容 超过1024时，倍扩容
//栈长度小的时候，节省空间，较大时，节省时间
func (s *Stack) Expansion()  {
	c := s.Cap
	if s.Cap <= 1024 {
		s.Cap = int(float64(s.Cap) * 1.5)
	}else{
		s.Cap = s.Cap << 1
	}
	newEleArr := make([]interface{},s.Cap)
	for i := 0; i < c; i++ {
		newEleArr[i] = s.elements[i]
	}
	s.elements = newEleArr
}

//入栈
func (s *Stack) Push(value interface{})  {
	if s.len == s.Cap {
		s.Expansion()
	}
	s.elements[s.len] = value
	s.len++
}

//判空
func (s *Stack) IsEmpty() bool {
	return s.len == 0
}

//获取长度
func (s *Stack) Len() int {
	return s.len
}

//出栈
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	ele := s.elements[s.len - 1]
	s.len--
	return ele
}

//获取栈顶元素
func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.elements[s.len - 1]
}
