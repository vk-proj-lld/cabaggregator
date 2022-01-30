package rider

import (
	"sync/atomic"
	"time"

	"github.com/vk-proj-lld/cabaggregator/entities"
)

var reqcounter uint32

type DriverSignal struct {
	sig      entities.AckSignal
	driverId int
}

func NewDriverSignal(sig entities.AckSignal, driverId int) DriverSignal {
	return DriverSignal{sig, driverId}
}

type RideRequest struct {
	id,
	riderId int

	sigchan chan<- DriverSignal
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

func (rr *RideRequest) RegisterSigChan(sigchan chan<- DriverSignal) {
	rr.sigchan = sigchan
}

func (rr *RideRequest) ReceiveSignal(sig DriverSignal) {
	if rr.sigchan != nil {
		rr.sigchan <- sig
	}
}
