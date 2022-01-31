package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/vk-proj-lld/cabaggregator/business/repo"
	"github.com/vk-proj-lld/cabaggregator/business/uc"
	"github.com/vk-proj-lld/cabaggregator/entities/driver"
	"github.com/vk-proj-lld/cabaggregator/entities/out"
	"github.com/vk-proj-lld/cabaggregator/entities/rider"
)

func main() {
	// testChanelBroadCasting()

	repo := repo.NewDispatcherRepo()
	stdo := out.NewStdO()
	log := out.NewFileOut()
	dispatcher := uc.NewDispatcher(repo, stdo, log)
	dispatcher.AddDriver(getDrivers(7)...)
	// fmt.Println(dispatcher.Dispatch(getRides(1)[0]))
	// return
	var wg sync.WaitGroup
	for _, ride := range getRides(4) {
		wg.Add(1)
		go func(ride *rider.RideRequest) {
			driver, err := dispatcher.Dispatch(ride)
			if err != nil {
				defer stdo.Write("no cab dound for ride:", ride.Id(), err)
			} else {
				defer stdo.Write("cab found for ride:", ride.Id(), driver)
			}
			wg.Done()
		}(ride)
	}
	wg.Wait()
	fmt.Println("Everything executed")
}

func getDrivers(n int) (drivers []*driver.Driver) {
	for i := 0; i < n; i++ {
		name := fmt.Sprintf("Driver-%d", i+1)
		drivers = append(drivers, driver.NewDriver(name, driver.NewEqualChoiceStrategy(driver.AckAccept, driver.AckReject)))
	}
	return drivers
}

func getRides(n int) (rides []*rider.RideRequest) {
	for i := 0; i < n; i++ {
		rides = append(rides, rider.NewRideRequest(i+1, time.Now()))
	}
	return rides
}
