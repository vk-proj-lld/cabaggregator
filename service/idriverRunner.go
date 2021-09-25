package service

import "ubermc/entity"

type idriverRunner interface {
	getDriver() *entity.Driver
	Accept(riderId string, respChan chan<- *entity.DriverResponse)
	isBusy() bool
	engage() bool
	disEngage() bool
}
