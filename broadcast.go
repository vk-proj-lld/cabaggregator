package main

import (
	"bufio"
	"fmt"
	"os"
	"sync/atomic"
)

func testChanelBroadCasting() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	msgchan := make(chan string)
	for i := 0; i < 3; i++ {
		NewRunner((chan string)(msgchan)).Run()
	}
L1:
	for {
		text, _ := reader.ReadString('\n')
		// fmt.Print(text)
		msgchan <- text
		if text == "quit" || text == "exit" {
			close(msgchan)
			break L1
		}
	}

}

var counter uint32

type runner struct {
	id   int
	data chan string
}

func NewRunner(data chan string) *runner {
	runner := &runner{
		id:   int(atomic.AddUint32(&counter, 1)),
		data: data,
	}
	return runner
}

func (r *runner) Run() {
	go func() {
	L1:
		for msg := range r.data {
			fmt.Printf("Runner:(%d) %s", r.id, msg)
			if msg == "quit" || msg == "exit" {
				close(r.data)
				break L1
			}
		}
	}()
}
