package idispatcher

import (
	"github.com/vk-proj-lld/cabaggregator/entities/driver"
	"github.com/vk-proj-lld/cabaggregator/entities/rider"
)

type IDispatcher interface {
	AddDriver(drivers ...*driver.Driver)
	Dispatch(ride *rider.RideRequest) (*driver.Driver, error)
}
