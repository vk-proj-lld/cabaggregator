package out

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"sync/atomic"
)

type fileout struct {
	mu     sync.Mutex
	writer io.WriteCloser
}

var counter uint32

func NewFileOut() IOout {
	w, err := os.OpenFile(
		fmt.Sprintf("./out.%d.txt", atomic.AddUint32(&counter, 1)),
		os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_TRUNC|os.O_SYNC,
		0644,
	)
	if err == nil {
		fmt.Println(w.Name())
	} else {
		fmt.Println(err)
	}
	return &fileout{
		writer: w,
	}
}

func (f *fileout) Write(data ...interface{}) {
	var raw []byte
	for _, datum := range data {
		if t, ok := datum.(interface{ String() string }); ok {
			raw = append(raw, []byte(t.String())...)
		} else if r, err := json.Marshal(datum); err != nil {
			raw = append(raw, r...)
		}
	}
	raw = append(raw, '\n')
	f.writer.Write(raw)
}

func (f *fileout) Close() error {
	return f.writer.Close()
}
