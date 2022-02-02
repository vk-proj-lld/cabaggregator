package out

import (
	"fmt"
	"sync"
)

type stdo struct {
	mu sync.Mutex
}

func NewStdO() IOout {
	return &stdo{}
}

func (out *stdo) Write(contents ...interface{}) {
	out.mu.Lock()
	defer out.mu.Unlock()
	fmt.Println("\n-----------------------------------")
	for _, content := range contents {
		fmt.Print(content, " ")
	}
}

func (out *stdo) Close() error {
	return nil
}
