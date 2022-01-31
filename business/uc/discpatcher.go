package uc

import (
	"errors"
	"sync"
	"time"

	"github.com/vk-proj-lld/cabaggregator/entities/driver"
	"github.com/vk-proj-lld/cabaggregator/entities/out"
	"github.com/vk-proj-lld/cabaggregator/entities/rider"
	"github.com/vk-proj-lld/cabaggregator/interfaces/idispatcher"
)

type rideDriverResp struct {
	*rider.RideRequest
	drchan chan<- *driver.Driver
}

type dispatcher struct {
	ridedrivers chan rideDriverResp
	disprepo    idispatcher.IDispatcherRepo

	out, logger out.IOout
}

func NewDispatcher(disprepo idispatcher.IDispatcherRepo, output, logout out.IOout) idispatcher.IDispatcher {
	if output == nil {
		output = out.NewFileOut()
	}
	if logout == nil {
		output = out.NewFileOut()
	}
	disp := &dispatcher{
		disprepo:    disprepo,
		ridedrivers: make(chan rideDriverResp),
		out:         output,
		logger:      logout,
	}
	go disp.run()
	return disp
}

/*
	request may get multiple AcceptSignal
	dispatcher have to block the available driver where it got the requests from,
	- one request can block multiple driver
	- driver should not be booked if blocked(serving riderRequest)
*/

//time taking function in real-world prolem, 5 mins.
func (disp *dispatcher) Dispatch(ride *rider.RideRequest) (*driver.Driver, error) {
	driverchan := make(chan *driver.Driver, 1)

	disp.ridedrivers <- rideDriverResp{ride, driverchan}
	driver := <-driverchan

	close(driverchan)
	if driver == nil {
		return nil, errors.New("no driver found")
	}
	return driver, nil
}

func (disp *dispatcher) AddDriver(drivers ...*driver.Driver) {
	for _, dr := range drivers {
		disp.disprepo.SaveDriver(dr)
	}
}

func (disp *dispatcher) run() {
	func() {
		for rd := range disp.ridedrivers {
			drivers := disp.disprepo.GetDrivers()
			disp.broadcast(rd.drchan, rd.RideRequest, drivers...)
		}
	}()
}

// first blockableDriver is to be returned
func (disp *dispatcher) broadcast(drchanel chan<- *driver.Driver, ride *rider.RideRequest, drivers ...*driver.Driver) {
	var unitmu sync.Once
	disp.broadcastHelper(&unitmu, drchanel, ride, drivers)
}

func (disp *dispatcher) broadcastHelper(unitmu *sync.Once, drchanel chan<- *driver.Driver, ride *rider.RideRequest, drivers []*driver.Driver) {
	var size = len(drivers)
	if size == 0 {
		return
	}
	if size == 1 {
		go disp.unicast(unitmu, drchanel, ride, drivers[0])
	}
	go disp.broadcastHelper(unitmu, drchanel, ride, drivers[0:size/2])  //  l1
	go disp.broadcastHelper(unitmu, drchanel, ride, drivers[size/2+1:]) // l2
}

func (disp *dispatcher) unicast(unitmu *sync.Once, drchanel chan<- *driver.Driver, ride *rider.RideRequest, drvr *driver.Driver) {
	if drvr.IsBlocked() {
		//for now blocked driver is not notified
		//can be discussed
		return
	}
	signal := drvr.Decide(ride, nil)
	disp.logger.Write(ride, drvr, signal, time.Now())
	if signal == driver.AckAccept {
		if drvr.Block(ride) {
			unitmu.Do(func() {
				drchanel <- drvr
			})
		}
	}
}
