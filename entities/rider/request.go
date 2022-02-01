package rider

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var reqcounter uint32

type RideRequest struct {
	id,
	riderId,
	driverId int
	mu    *sync.Mutex
	wg    *sync.WaitGroup
	rtime time.Time
}

func NewRideRequest(riderId int, reqtime time.Time) *RideRequest {
	return &RideRequest{
		id:      int(atomic.AddUint32(&reqcounter, 1)),
		riderId: riderId,
		rtime:   reqtime,
		mu:      &sync.Mutex{},
		wg:      &sync.WaitGroup{},
	}
}

func (rr *RideRequest) Id() int { return rr.id }

func (rr *RideRequest) RiderId() int { return rr.riderId }

func (rr *RideRequest) String() string {
	return fmt.Sprintf("Rider (%d) with requestId (%d) at %s", rr.riderId, rr.id, rr.rtime.Format("2006-01-02 15:04:05"))
}

func (rr *RideRequest) DriverId() int {
	return rr.driverId
}

func (rr *RideRequest) SetDriverIfNotSet(id int) bool {
	if rr.driverId != 0 {
		return false
	}
	rr.mu.Lock()
	defer rr.mu.Unlock()
	if rr.driverId != 0 {
		return false
	}
	rr.driverId = id
	return true
}

func (rr *RideRequest) GetWG() *sync.WaitGroup { return rr.wg }
