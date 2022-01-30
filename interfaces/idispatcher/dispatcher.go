package idispatcher

import "github.com/vk-proj-lld/cabdispatcher/entities"

type IDispatcher interface {
	AddDriver(ids ...string)
	Dispatch(*entities.Ride) (*entities.Driver, error)
}
