package driver

import (
	"fmt"
)

type DriverSignal struct {
	sig              AckSignal
	driverId, rideId int
}

func NewDriverSignal(sig AckSignal, driverId, rideId int) DriverSignal {
	return DriverSignal{sig: sig, driverId: driverId, rideId: rideId}
}

func (dsig DriverSignal) Sig() AckSignal {
	return dsig.sig
}

func (dsig DriverSignal) DriverId() int {
	return dsig.driverId
}

func (dsig DriverSignal) String() string {
	return fmt.Sprintf("Driver id(%d) responds with %v to Rider Id(%d)", dsig.driverId, dsig.sig.String(), dsig.rideId)
}
