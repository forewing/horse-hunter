package resources

import (
	_ "embed"
	"encoding/base64"
	"encoding/json"
)

var (
	//go:embed max.json
	embedMax []byte
	//go:embed min.json
	embedMin []byte
	//go:embed replace.json
	embedReplace []byte

	LinesMax     []string
	LinesMin     []string
	LinesReplace map[string]string
)

func init() {
	err := loadDecodeSlice(embedMax, &LinesMax)
	if err != nil {
		panic(err)
	}

	err = loadDecodeSlice(embedMin, &LinesMin)
	if err != nil {
		panic(err)
	}

	linesReplaceOrigin := map[string]string{}
	err = json.Unmarshal(embedReplace, &linesReplaceOrigin)
	if err != nil {
		panic(err)
	}

	LinesReplace = make(map[string]string)
	for k, v := range linesReplaceOrigin {
		k2, err := decode(k)
		if err != nil {
			panic(err)
		}
		v2, err := decode(v)
		if err != nil {
			panic(err)
		}
		LinesReplace[k2] = v2
	}
}

func loadDecodeSlice(data []byte, target *[]string) (err error) {
	err = json.Unmarshal(data, target)
	if err != nil {
		return err
	}
	for i := range *target {
		(*target)[i], err = decode((*target)[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func decode(s string) (string, error) {
	d, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(d), nil
}
