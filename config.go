package hunter

import (
	"time"
)

// Level of insult
type Level int

// Target of insult
type Target int

const (
	LevelMax     Level = 0
	LevelMin     Level = 1
	LevelMix     Level = 2
	LevelDefault Level = LevelMax

	TargetFemale  Target = 0
	TargetMale    Target = 1
	TargetMix     Target = 2
	TargetDefault Target = TargetFemale

	IntervalDefault time.Duration = time.Millisecond * 200
)

var (
	LevelLookup = map[string]Level{
		"max": LevelMax,
		"min": LevelMin,
		"mix": LevelMix,
	}

	TargetLookup = map[string]Target{
		"female": TargetFemale,
		"male":   TargetMale,
		"mix":    TargetMix,
	}

	LevelName  = []string{"max", "min", "mix"}
	TargetName = []string{"female", "male", "mix"}
)
