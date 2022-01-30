package strategy

import (
	"math/rand"
	"time"

	"github.com/vk-proj-lld/cabdispatcher/interfaces/istrategy"
	"github.com/vk-proj-lld/cabdispatcher/utils"
)

type equalChoiceStrategy struct {
	choices []string
	slots   int
	rangen  *rand.Rand
}

func NewEqualChoiceStrategy(choices ...string) istrategy.IStrategy {
	return &equalChoiceStrategy{
		slots:   len(choices),
		choices: choices,
		rangen:  rand.New(rand.NewSource(utils.RandomGenSeed)),
	}
}

func (eqst *equalChoiceStrategy) Select() string {
	//processing time
	ms := 100 + eqst.rangen.Intn(500)
	time.Sleep(time.Duration(ms))

	return eqst.choices[eqst.rangen.Intn(eqst.slots)]
}
