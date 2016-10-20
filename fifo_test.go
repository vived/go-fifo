package go_fifo

import (
	"fmt"
	"testing"
	"time"
)

func custom(q *Queue, id int, chexit chan int) {

	timeout := false

	for !timeout {
		select {

		case <-time.After(time.Second * 5):
			timeout = true
			break
		case v := <-q.GetChannel():
			fmt.Printf("------id:%d custom:%d\n", id, v)
			time.Sleep(time.Second)
		}

		if timeout {
			break
		}

	}

	chexit <- 0
}

func product(q *Queue, id int, max int) {

	for {
		count++
		fmt.Printf("id:%d product:%d\n", id, count)
		q.Put(count)
		if count >= max {
			break
		}
	}
}

var count = 0

func TestFifo(t *testing.T) {
	q := NewFifoQueue(10)
	threadcount := 8
	chexit := make([]chan int, threadcount)
	for i := 0; i < threadcount; i++ {
		chexit[i] = make(chan int)
		go custom(q, i, chexit[i])
	}

	for i := 0; i < 1; i++ {
		go product(q, i, 20)
	}

	for i := 0; i < threadcount; i++ {
		<-chexit[i]
	}

}
