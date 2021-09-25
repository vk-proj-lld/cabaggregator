package main

import (
	"fmt"
	"sync"
	"ubermc/entity"
	"ubermc/service"
)

func main() {

	var dispather service.IDispatcher = service.NewDispatcher()

	dispather.AddDriver("D1", "D2", "D3", "D4", "D5")

	var wg sync.WaitGroup
	var rides []*entity.Ride
	for i := 0; i < 10; i++ {
		rides = append(rides, entity.NewRide(fmt.Sprintf("R%d", i+1)))
	}
	for _, ride := range rides {
		wg.Add(1)
		go func(ride *entity.Ride) {
			driver, err := dispather.Dispatch(ride)
			if err != nil {
				fmt.Printf("Booking Ride for %s, No drivers\n", ride.GetId())
			} else {
				fmt.Printf("Booking Ride for %s,Driver assigned:%s\n", ride.GetId(), driver.GetId())
			}
			wg.Done()
		}(ride)
	}
	wg.Wait()
}
