package go_queue

type Queue struct {
	list *queueFifo
}

type queueFifo struct {
	count   uint16
	chvalue chan interface{}
}

//通过此函数创建队列
func NewFifoQueue(count uint16) *Queue {
	qfifo := &queueFifo{count, make(chan interface{}, count)}

	q := &Queue{}
	q.list = qfifo
	return q
}

func (q *Queue) Put(obj interface{}) {

	q.list.chvalue <- obj
}

func (q *Queue) Get() interface{} {
	obj := <-q.list.chvalue
	return obj
}

func (q *Queue) GetChannel() chan interface{} {
	return q.list.chvalue
}
