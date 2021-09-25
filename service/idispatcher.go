package service

import "ubermc/entity"

type IDispatcher interface {
	AddDriver(ids ...string)
	Dispatch(*entity.Ride) (*entity.Driver, error)
}
