package entity

import "sync"

type Ride struct {
	id string

	booked bool
	mu     sync.Mutex
}

func NewRide(id string) *Ride {
	return &Ride{id: id}
}

func (r *Ride) GetId() string { return r.id }

func (r *Ride) Book() bool {
	if !r.booked {
		r.mu.Lock()
		defer r.mu.Unlock()
		if !r.booked {
			r.booked = true
			return true
		}
	}
	return false
}
