package rider

import (
	"fmt"
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

func (rr *RideRequest) String() string {
	return fmt.Sprintf("Rider (%d) with requestId (%d) at %T", rr.riderId, rr.id, rr.rtime)
}
