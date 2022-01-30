package uc

import "github.com/vk-proj-lld/cabaggregator/entities/rider"

type dispatcher struct {
	rides chan *rider.RideRequest
}
