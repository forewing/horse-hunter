package hunter

// Level of insult
type Level int

// Target of insult
type Target int

const (
	LevelMax Level = 0
	LevelMin Level = 1
	LevelMix Level = 2

	TargetFemale Target = 0
	TargetMale   Target = 1
	TargetMix    Target = 2
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
)
