package istrategy

import "github.com/vk-proj-lld/cabaggregator/entities"

type IStrategy interface {
	Select() entities.AckSignal
}
