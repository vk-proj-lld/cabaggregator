package uc

import (
	"errors"
	"time"

	"github.com/vk-proj-lld/cabaggregator/entities/driver"
	"github.com/vk-proj-lld/cabaggregator/entities/out"
	"github.com/vk-proj-lld/cabaggregator/entities/rider"
	"github.com/vk-proj-lld/cabaggregator/interfaces/idispatcher"
)

type dispatcher struct {
	rides       chan *rider.RideRequest
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
		disprepo: disprepo,
		rides:    make(chan *rider.RideRequest),
		out:      output,
		logger:   logout,
	}
	go disp.run()
	return disp
}

func (disp *dispatcher) AddDriver(drivers ...*driver.Driver) {
	for _, dr := range drivers {
		disp.disprepo.SaveDriver(dr)
	}
}

/*
	request may get multiple AcceptSignal
	dispatcher have to block the available driver where it got the requests from,
	- one request can block multiple driver
	- driver should not be booked if blocked(serving riderRequest)
*/

//time taking function in real-world problem, 5 mins.
func (disp *dispatcher) Dispatch(ride *rider.RideRequest) (*driver.Driver, error) {
	ride.GetWG().Add(1)
	disp.rides <- ride
	ride.GetWG().Wait()
	if ride.DriverId() == 0 {
		return nil, errors.New("no driver found")
	}
	return disp.disprepo.GetDriver(ride.DriverId()), nil
}

func (disp *dispatcher) run() {
	func() {
		for ride := range disp.rides {
			drivers := disp.disprepo.GetDrivers()
			ride.GetWG().Add(len(drivers))
			disp.broadcast(ride, drivers)
			ride.GetWG().Done()
		}
	}()
}

// first blockableDriver is to be returned
func (disp *dispatcher) broadcast(ride *rider.RideRequest, drivers []*driver.Driver) {
	var size = len(drivers)
	if size == 0 {
		return
	}
	if size == 1 {
		go disp.unicast(ride, drivers[0])
		return
	}
	go disp.broadcast(ride, drivers[:size/2])
	go disp.broadcast(ride, drivers[size/2:])
}

func (disp *dispatcher) unicast(ride *rider.RideRequest, drvr *driver.Driver) {
	defer ride.GetWG().Done()
	if drvr.IsBlocked() {
		//for now blocked driver is not notified
		//can be discussed
		return
	}
	signal := drvr.Decide(ride, nil)
	if signal == driver.AckAccept && drvr.Block(ride) {
		disp.logger.Write(ride, drvr, signal, "Booked", time.Now().Format("2006-01-02 15:04:05"))
	} else {
		disp.logger.Write(ride, drvr, signal, "Not Booked", time.Now().Format("2006-01-02 15:04:05"))
	}
}
