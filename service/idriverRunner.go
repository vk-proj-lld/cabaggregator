package service

import "github.com/vk-proj-lld/cabdispatcher/entities"

type idriverRunner interface {
	getDriver() *entities.Driver
	Accept(riderId string, respChan chan<- *entities.DriverResponse)
	isBusy() bool
	engage() bool
	disEngage() bool
}
