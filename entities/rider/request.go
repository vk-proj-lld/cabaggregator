package rider

import (
	"sync/atomic"
	"time"
)

var reqcounter uint32

type RideRequest struct {
	id,
	riderId int

	rtime time.Time
}

func NewRideRequest(riderId int, reqtime time.Time) *RideRequest {
	return &RideRequest{
		id:      int(atomic.AddUint32(&reqcounter, 1)),
		riderId: riderId,
		rtime:   reqtime,
	}
}

func (rr *RideRequest) Id() int { return rr.id }

func (rr *RideRequest) RiderId() int { return rr.riderId }
