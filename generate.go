package hunter

import (
	"math/rand"
	"strings"
	"time"

	"github.com/forewing/horse-hunter/resources"
)

var (
	rng *rand.Rand
)

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	rng = rand.New(source)
}

// GetLine by level and target arguments
func GetLine(level Level, target Target) string {
	lines := &resources.LinesMax
	if level == LevelMin || (level == LevelMix && rng.Float64() < 0.5) {
		lines = &resources.LinesMin
	}

	line := (*lines)[rng.Intn(len(*lines))]
	if target == TargetMale || (target == TargetMix && rng.Float64() < 0.5) {
		for k, v := range resources.LinesReplace {
			line = strings.ReplaceAll(line, k, v)
		}
	}

	return line
}
