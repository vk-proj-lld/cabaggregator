package istrategy

import "github.com/vk-proj-lld/cabaggregator/entities/signals"

type IStrategy interface {
	Select() signals.AckSignal
}
