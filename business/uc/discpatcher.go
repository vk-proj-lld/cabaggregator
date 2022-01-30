package uc

import (
	"errors"
	"fmt"

	"github.com/vk-proj-lld/cabaggregator/entities/driver"
	"github.com/vk-proj-lld/cabaggregator/entities/rider"
	"github.com/vk-proj-lld/cabaggregator/interfaces/idispatcher"
)

type dispatcher struct {
	rides    chan *rider.RideRequest
	disprepo idispatcher.IDispatcherRepo
}

func NewDispatcher() idispatcher.IDispatcher {
	disp := &dispatcher{}
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
	sigchan := make(chan rider.DriverSignal)
	ride.RegisterSigChan(sigchan)
	disp.rides <- ride
	return disp.listenNBlockAvailableDriver(sigchan)
}

func (disp *dispatcher) AddDriver(ids ...string) {

}

func (disp *dispatcher) run() {
	func() {
		for ride := range disp.rides {
			drivers := disp.disprepo.GetDrivers()
			disp.broadcast(ride, drivers...)
		}
	}()
}

func (disp *dispatcher) broadcast(ride *rider.RideRequest, drivers ...*driver.Driver) {
	for _, drvr := range drivers {
		go drvr.InformIncommingRide(ride)
	}
}

// in the given channels all the signal from various drivers are listened
// first blockableDriver is to be returned
func (disp *dispatcher) listenNBlockAvailableDriver(sigs <-chan rider.DriverSignal) (*driver.Driver, error) {
	for sig := range sigs {
		fmt.Println(sig)
	}
	return nil, errors.New("no driver found")
}
