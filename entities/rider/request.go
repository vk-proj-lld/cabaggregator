package rider

import (
	"sync/atomic"
	"time"

	"github.com/vk-proj-lld/cabaggregator/entities/driver"
)

var reqcounter uint32

type RideRequest struct {
	id,
	riderId int

	dchan   chan<- *driver.Driver
	sigchan chan<- driver.DriverSignal
	rtime   time.Time
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

func (rr *RideRequest) RegisterDriverChan(dchan chan<- *driver.Driver) {
	rr.dchan = dchan
}

func (rr *RideRequest) GetDriverChan() chan<- *driver.Driver {
	return rr.dchan
}

func (rr *RideRequest) RegisterSigChan(sigchan chan<- driver.DriverSignal) {
	rr.sigchan = sigchan
}

func (rr *RideRequest) GetSigChan() chan<- driver.DriverSignal {
	return rr.sigchan
}

func (rr *RideRequest) ReceiveSignal(sig driver.DriverSignal) {
	if rr.sigchan != nil {
		rr.sigchan <- sig
	}
}
