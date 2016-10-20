package go_queue

import (
	"fmt"
	"testing"
	"time"
)

func custom(q *Queue, id int, chexit chan int) {

	timeout := false
	max := 0

	for !timeout {
		select {

		case <-time.After(time.Microsecond * 1000):
			timeout = true
			break
		case v := <-q.GetChannel():
			fmt.Printf("------id:%d custom:%d\n", id, v)
			max += v.(int)
			time.Sleep(time.Second)
		}

		if timeout {
			break
		}

	}

	chexit <- max
}

func product(q *Queue, id int, max int) {

	for {
		count++
		fmt.Printf("id:%d product:%d\n", id, count)
		q.Put(count)
		sum += count
		if count >= max {
			break
		}
	}
}

var count = 0
var sum = 0

func TestFifo(t *testing.T) {
	q := NewFifoQueue(10)
	threadcount := 4
	chexit := make([]chan int, threadcount)
	for i := 0; i < threadcount; i++ {
		chexit[i] = make(chan int)
		go custom(q, i, chexit[i])
	}

	for i := 0; i < 1; i++ {
		go product(q, i, 22)
	}

	var custommax = 0
	for i := 0; i < threadcount; i++ {
		custommax += <-chexit[i]
	}

	if sum != custommax {
		t.Errorf("process:%d not equal to custom:%d", sum, custommax)
	}
	t.Logf("process:%d is equal to custom:%d", sum, custommax)

}
