package entity

import "sync"

type Driver struct {
	id string
	mu sync.Mutex

	blocked bool
}

func NewDriver(id string) *Driver {
	return &Driver{id: id}
}

func (d *Driver) GetId() string { return d.id }

func (d *Driver) Block() bool {
	if !d.blocked {
		d.mu.Lock()
		defer d.mu.Unlock()
		if !d.blocked {
			d.blocked = true
			return true
		}
	}
	return false
}
