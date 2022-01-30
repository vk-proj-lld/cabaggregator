package strategy

import (
	"math/rand"
	"time"

	"github.com/vk-proj-lld/cabaggregator/entities"
	"github.com/vk-proj-lld/cabaggregator/interfaces/istrategy"
	"github.com/vk-proj-lld/cabaggregator/utils"
)

type equalChoiceStrategy struct {
	choices []entities.AckSignal
	slots   int
	rangen  *rand.Rand
}

func NewEqualChoiceStrategy(choices ...entities.AckSignal) istrategy.IStrategy {
	return &equalChoiceStrategy{
		slots:   len(choices),
		choices: choices,
		rangen:  rand.New(rand.NewSource(utils.RandomGenSeed)),
	}
}

func (eqst *equalChoiceStrategy) Select() entities.AckSignal {
	//processing time
	ms := 100 + eqst.rangen.Intn(500)
	time.Sleep(time.Duration(ms))

	return eqst.choices[eqst.rangen.Intn(eqst.slots)]
}
