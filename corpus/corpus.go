package corpus

import (
	_ "embed"
	"encoding/json"
)

//go:embed english_1k.json
var english1k []byte

func Enghlish1k() []string {
	var words []string

	if err := json.Unmarshal(english1k, &words); err != nil {
		panic("failed to parse english_1k.json")
	}

	return words
}
