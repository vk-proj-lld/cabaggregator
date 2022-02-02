package driver

import (
	"math/rand"
	"time"

	"github.com/vk-proj-lld/cabaggregator/utils"
)

type equalChoiceStrategy struct {
	choices []AckSignal
	slots   int
	mintime time.Duration
}

var rangen = rand.New(rand.NewSource(utils.RandomGenSeed))

func NewEqualChoiceStrategy(mintime time.Duration, choices ...AckSignal) IStrategy {
	return &equalChoiceStrategy{
		mintime: mintime,
		slots:   len(choices),
		choices: choices,
	}
}

func (eqst *equalChoiceStrategy) Select() AckSignal {
	//processing time
	ms := eqst.mintime + time.Millisecond*(time.Duration(rangen.Intn(500)))
	time.Sleep(ms)

	idx := rangen.Intn(eqst.slots)
	return eqst.choices[idx]
}
