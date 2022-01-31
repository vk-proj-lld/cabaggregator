package driver

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter uint32

type Driver struct {
	id   int
	name string

	choiceStrategy IStrategy

	mu      sync.Mutex
	blocked bool
}

func NewDriver(name string, strat IStrategy) *Driver {
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

func (d *Driver) InformIncommingRide(rideId int, dsig chan<- DriverSignal) {
	dsig <- NewDriverSignal(d.choiceStrategy.Select(), d.id, rideId)
}

func (d *Driver) Block() bool {
	if d.blocked {
		return false
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.blocked {
		return false
	} else {
		d.blocked = true
		return true
	}
}
