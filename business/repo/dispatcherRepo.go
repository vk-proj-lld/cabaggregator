package repo

import (
	"github.com/vk-proj-lld/cabaggregator/entities/driver"
)

type disprepo struct {
	drivers map[int]*driver.Driver
}

func NewDispatcherRepo() *disprepo {
	return &disprepo{
		drivers: make(map[int]*driver.Driver),
	}
}

func (d *disprepo) SaveDriver(drvr *driver.Driver) error {
	d.drivers[drvr.Id()] = drvr
	return nil
}

func (d *disprepo) GetDriver(driverId int) *driver.Driver { return d.drivers[driverId] }

func (d *disprepo) GetDrivers() (drivers []*driver.Driver) {
	for _, drvr := range d.drivers {
		drivers = append(drivers, drvr)
	}
	return drivers
}
