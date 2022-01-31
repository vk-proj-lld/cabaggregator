package uc

import (
	"errors"
	"sync"

	"github.com/vk-proj-lld/cabaggregator/entities/driver"
	"github.com/vk-proj-lld/cabaggregator/entities/out"
	"github.com/vk-proj-lld/cabaggregator/entities/rider"
	"github.com/vk-proj-lld/cabaggregator/interfaces/idispatcher"
	"github.com/vk-proj-lld/cabaggregator/interfaces/iio"
)

type rideDriverResp struct {
	*rider.RideRequest
	drchan chan<- *driver.Driver
}

type dispatcher struct {
	ridedrivers chan rideDriverResp
	disprepo    idispatcher.IDispatcherRepo

	out iio.IOout
}

func NewDispatcher(disprepo idispatcher.IDispatcherRepo, output iio.IOout) idispatcher.IDispatcher {
	if output == nil {
		output = out.NewConsoleOutPutUsecase()
	}
	disp := &dispatcher{
		disprepo:    disprepo,
		ridedrivers: make(chan rideDriverResp),
		out:         output,
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

func (disp *dispatcher) Dispatch(ride *rider.RideRequest) (*driver.Driver, error) {
	driverchan := make(chan *driver.Driver, 1)
	defer close(driverchan)

	disp.ridedrivers <- rideDriverResp{ride, driverchan}

	driver := <-driverchan
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
	var mu sync.Mutex
	var driverAssigned bool
	for _, drvr := range drivers {
		go func(drvr *driver.Driver) {
			signal := drvr.Decide(ride, nil)
			if signal == driver.AckAccept && !driverAssigned {
				mu.Lock()
				defer mu.Unlock()
				if !driverAssigned && drvr.Block(ride) {
					driverAssigned = true
					drchanel <- drvr
				}
			}
		}(drvr)
	}
}
