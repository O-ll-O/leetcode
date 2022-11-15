package queue

type Queue interface {
	Enqueue(data int) bool
	Dequeue() int
	IsEmpty() bool
}
type queue struct {
	datas []int
	n     int
	head  int
	tail  int
}

func NewQueue(capacity int) Queue {
	return &queue{
		datas: make([]int, capacity),
		n:     capacity,
	}
}

func (q *queue) Enqueue(data int) bool {
	if q.tail == q.n {
		return false
	}
	q.datas[q.tail] = data
	q.tail++
	return true
}

func (q *queue) Dequeue() int {
	if q.head == q.tail {
		return -1
	}
	v := q.datas[q.head]
	q.head++
	return v
}

func (q *queue) IsEmpty() bool {
	return q.head == q.tail
}
