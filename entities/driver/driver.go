package driver

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/vk-proj-lld/cabaggregator/entities/rider"
	"github.com/vk-proj-lld/cabaggregator/interfaces/istrategy"
)

var counter uint32

type Driver struct {
	id   int
	name string

	choiceStrategy istrategy.IStrategy

	stop chan struct{}

	mu      sync.Mutex
	blocked bool
}

func NewDriver(name string, strat istrategy.IStrategy) *Driver {
	return &Driver{
		id:             int(atomic.AddUint32(&counter, 1)),
		name:           name,
		choiceStrategy: strat,
	}
}

func (d *Driver) Id() int { return d.id }

func (d *Driver) Name() string { return d.name }

func (d *Driver) String() string {
	return fmt.Sprintf("(%d) %s", d.id, d.name)
}

func (d *Driver) RegisterRideChannel(rides <-chan *rider.RideRequest) {
	go func() {
	L1:
		for {
			select {
			case ride := <-rides:
				fmt.Println(ride)
				//todo -  take decision take it or now
			case <-d.stop:
				break L1
			}
		}
	}()
}

func (d *Driver) DeRegisterRideChannel() {
	d.stop <- struct{}{}
}
