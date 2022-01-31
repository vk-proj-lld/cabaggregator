package rider

import (
	"fmt"
	"sync/atomic"
)

var counter uint32

type Rider struct {
	id   int
	name string
}

func NewRider(name string) *Rider {
	return &Rider{
		id:   int(atomic.AddUint32(&counter, 1)),
		name: name,
	}
}

func (r *Rider) Id() int { return r.id }

func (r *Rider) Name() string { return r.name }

func (r *Rider) String() string { return fmt.Sprintf("%s(%d)", r.name, r.id) }
