package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/vk-proj-lld/cabdispatcher/entities"
)

type dispather struct {
	drivers map[string]idriverRunner
	// acticeDrivers []idriverRunner
}

func NewDispatcher() IDispatcher {
	return &dispather{drivers: map[string]idriverRunner{}}
}

func (disp *dispather) RemoveDriver(id string) {

}

func (disp *dispather) AddDriver(ids ...string) {
	for _, id := range ids {
		disp.drivers[id] = NewDriverRunner(id)
	}
}

func (disp *dispather) Dispatch(ride *entities.Ride) (*entities.Driver, error) {
	var responseChan = make(chan *entities.DriverResponse)
	var foundDrunner idriverRunner

	var wg sync.WaitGroup
	var broadcasted = 0
	for _, drunner := range disp.drivers {
		if !drunner.isBusy() {
			broadcasted++
			wg.Add(1)
			go drunner.Accept(ride.GetId(), responseChan)
		}
	}
	var mu sync.Mutex
	for i := 0; i < broadcasted; i++ {
		resp := <-responseChan
		go func() {
			fmt.Println(resp)
			if resp.GetMessage() == ActionAccept {
				mu.Lock()
				isEngaged := disp.drivers[resp.GetDriverId()].engage()
				if isEngaged {
					status := ride.Book()
					if status {
						foundDrunner = disp.drivers[resp.GetDriverId()]
					} else {
						disp.drivers[resp.GetDriverId()].disEngage()
					}
				}
				mu.Unlock()
			}
			wg.Done()
		}()
	}
	close(responseChan)
	wg.Wait()
	if foundDrunner == nil {
		return nil, errors.New("no driver found")
	}
	return foundDrunner.getDriver(), nil
}

// func(disp *)
