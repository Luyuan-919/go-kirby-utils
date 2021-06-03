package DataStruct

const (
	DefCapOfQueue = 32
)

type Queue struct {
	elements []interface{}
	len int
	cap int
	head int
	tail int
}

//新建一个队列
func NewQueue() *Queue {
	return &Queue{
		elements: make([]interface{},DefCapOfQueue),
		len: 0,
		cap: DefCapOfQueue,
		head: 0,
		tail: 0,
	}
}

//新建一个队列 自定义容量
func NewQueueWithCap(cap int) *Queue {
	if cap <= 0 {
		cap = DefCapOfQueue
	}
	return &Queue{
		elements: make([]interface{},DefCapOfQueue),
		len: 0,
		cap: cap,
		head: 0,
		tail: 0,
	}
}

//入队
func (q *Queue) Push(value interface{}) {
	if q.tail == q.cap {
		q.Expansion()
	}
	q.elements[q.tail] = value
	q.tail++
	q.len++
}

func (q *Queue) IsEmpty() bool {
	return q.len == 0
}

//出队
func (q *Queue) Pop() interface{} {
	if q.IsEmpty() {
		return nil
	}
	ele := q.elements[q.head]
	q.head++
	q.len--
	return ele
}

//获取队首元素
func (q *Queue) Head() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.elements[q.head]
}

//获取队尾元素
func (q *Queue) Tail() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.elements[q.tail - 1]
}

//获取队列长度
func (q *Queue) Len() int {
	return q.len
}

//扩容
func (q *Queue) Expansion() {
	if q.cap <= 1024 {
		q.cap = int(float64(q.cap) * 1.5)
	}else{
		q.cap = q.cap << 1
	}
	newEleArr := make([]interface{},q.cap)
	I := 0
	for i := q.head; i < q.tail; i++ {
		newEleArr[I] = q.elements[i]
		I++
	}
	q.elements = newEleArr
	q.head = 0
	q.tail = q.len
}