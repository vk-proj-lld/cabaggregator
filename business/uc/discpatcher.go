package uc

import (
	"errors"

	"github.com/vk-proj-lld/cabaggregator/entities/driver"
	"github.com/vk-proj-lld/cabaggregator/entities/out"
	"github.com/vk-proj-lld/cabaggregator/entities/rider"
	"github.com/vk-proj-lld/cabaggregator/interfaces/idispatcher"
	"github.com/vk-proj-lld/cabaggregator/interfaces/iio"
)

type dispatcher struct {
	rides    chan *rider.RideRequest
	disprepo idispatcher.IDispatcherRepo

	out iio.IOout
}

func NewDispatcher(disprepo idispatcher.IDispatcherRepo, output iio.IOout) idispatcher.IDispatcher {
	if output == nil {
		output = out.NewConsoleOutPutUsecase()
	}
	disp := &dispatcher{
		disprepo: disprepo,
		rides:    make(chan *rider.RideRequest),
		out:      output,
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
	ride.RegisterDriverChan(driverchan)

	disp.rides <- ride

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
		for ride := range disp.rides {
			drivers := disp.disprepo.GetDrivers()
			sigchan := make(chan driver.DriverSignal, len(drivers))
			ride.RegisterSigChan(sigchan)

			go disp.listenSignalsFromDriversAgainstRideRequest(ride, sigchan)
			disp.broadcast(ride, drivers...)
		}
	}()
}

func (disp *dispatcher) broadcast(ride *rider.RideRequest, drivers ...*driver.Driver) {
	for _, drvr := range drivers {
		go drvr.InformIncommingRide(ride.Id(), ride.GetSigChan())
	}
}

func (disp *dispatcher) listenSignalsFromDriversAgainstRideRequest(ride *rider.RideRequest, sigchan chan driver.DriverSignal) {
L1:
	for sig := range sigchan {
		if sig.Sig() == driver.AckAccept {
			driver := disp.disprepo.GetDriver(sig.DriverId())
			if driver != nil && driver.Block() {
				ride.GetDriverChan() <- driver
				break L1
			}
		}
		disp.out.Write(sig)
	}
	for sig := range sigchan {
		disp.out.Write(sig)
	}

	close(sigchan)
}

// in the given channels all the signal from various drivers are listened
// first blockableDriver is to be returned
