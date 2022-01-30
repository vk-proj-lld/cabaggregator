package out

import (
	"fmt"
	"sync"

	"github.com/vk-proj-lld/cabaggregator/interfaces/iio"
)

type outUseCase struct {
	mu sync.Mutex
}

func NewConsoleOutPutUsecase() iio.IOout {
	return &outUseCase{}
}

func (out *outUseCase) Write(contents ...interface{}) {
	out.mu.Lock()
	defer out.mu.Unlock()
	fmt.Println("-----------------------------------")
	for _, content := range contents {
		fmt.Println(content)
	}
	fmt.Println("-----------------------------------")
}
