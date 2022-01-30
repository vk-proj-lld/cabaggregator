package uc

import "github.com/vk-proj-lld/cabdispatcher/entities/rider"

type dispatcher struct {
	rides chan *rider.RideRequest
}
