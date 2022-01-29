package entities

import "fmt"

type DriverResponse struct {
	message, id, rideId string
}

func NewDriverResponse(message, id, rideId string) *DriverResponse {
	return &DriverResponse{message: message, id: id, rideId: rideId}
}

func (dresp *DriverResponse) String() string {
	return fmt.Sprintf("Driver %s's response on ride %s is %s", dresp.id, dresp.rideId, dresp.message)
}

func (dresp *DriverResponse) GetMessage() string { return dresp.message }

func (dresp *DriverResponse) GetDriverId() string { return dresp.id }
