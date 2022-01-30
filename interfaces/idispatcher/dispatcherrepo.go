package idispatcher

import (
	"github.com/vk-proj-lld/cabaggregator/entities/driver"
)

type IDispatcherRepo interface {
	SaveDriver(drvr *driver.Driver) error
	GetDriver(driverId int) *driver.Driver
	GetDrivers() []*driver.Driver
}
