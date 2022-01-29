package service

import (
	"math/rand"
	"sync"
	"time"

	"github.com/vk-proj-lld/cabdispatcher/entities"
)

const (
	maxProcessingTime = time.Second * 5
	ActionReject      = "rejecting"
	ActionAccept      = "accepting"
)

type driverRunner struct {
	processingTime time.Duration
	rangen         *rand.Rand
	driver         *entities.Driver

	mu   sync.Mutex
	busy bool
}

func NewDriverRunner(id string) idriverRunner {
	return &driverRunner{
		rangen:         rand.New(rand.NewSource(time.Now().UnixMicro())),
		driver:         entities.NewDriver(id),
		processingTime: maxProcessingTime,
	}
}
func (drunner *driverRunner) process() {
	//maybe the driver wants to do while working on the ride request.
	time.Sleep(time.Duration(drunner.rangen.Intn(5000)) * time.Millisecond)
}

func (drunner *driverRunner) Accept(rideId string, respChan chan<- *entities.DriverResponse) {
	if drunner.isBusy() {
		respChan <- entities.NewDriverResponse(ActionReject, drunner.driver.GetId(), rideId)
		return
	}
	drunner.process()
	if drunner.rangen.Intn(2) == 0 {
		respChan <- entities.NewDriverResponse(ActionAccept, drunner.driver.GetId(), rideId)
	} else {
		respChan <- entities.NewDriverResponse(ActionReject, drunner.driver.GetId(), rideId)
	}
}
func (drunner *driverRunner) getDriver() *entities.Driver { return drunner.driver }

func (drunner *driverRunner) change(state bool) bool {
	if drunner.busy != state {
		drunner.mu.Lock()
		defer drunner.mu.Unlock()
		if drunner.busy != state {
			drunner.busy = state
			return true
		}
	}
	return false
}

func (drunner *driverRunner) isBusy() bool { return drunner.busy }

func (drunner *driverRunner) engage() bool {
	return drunner.change(true)
}

func (drunner *driverRunner) disEngage() bool {
	return drunner.change(false)
}
