package strategy

import (
	"math/rand"
	"time"

	"github.com/vk-proj-lld/cabaggregator/entities/signals"
	"github.com/vk-proj-lld/cabaggregator/interfaces/istrategy"
	"github.com/vk-proj-lld/cabaggregator/utils"
)

type equalChoiceStrategy struct {
	choices []signals.AckSignal
	slots   int
}

var rangen = rand.New(rand.NewSource(utils.RandomGenSeed))

func NewEqualChoiceStrategy(choices ...signals.AckSignal) istrategy.IStrategy {
	return &equalChoiceStrategy{
		slots:   len(choices),
		choices: choices,
	}
}

func (eqst *equalChoiceStrategy) Select() signals.AckSignal {
	//processing time
	ms := 100 + rangen.Intn(500)
	time.Sleep(time.Duration(ms))

	idx := rangen.Intn(eqst.slots)
	return eqst.choices[idx]
}
