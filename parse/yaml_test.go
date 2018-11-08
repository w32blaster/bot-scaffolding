package parse

import (
	"fmt"
	"testing"
)

func TestParseFile(t *testing.T) {

	root, _ := Parse("./demo.bot.yaml")

	for k, v := range root.Commands {
		fmt.Printf(" - %s \n", k)

		for kk, vv := range v.Message {
			fmt.Printf("    message:  %s:\n       %s \n", kk, vv)
		}

		for i, row := range v.ButtonRows {
			fmt.Printf("     ROW:: %d \n", i)
			for _, button := range row.Buttons {
				fmt.Printf("     action: %s: \n label: %s \n", button.Action, button.Label)
			}
		}
	}
}

/*
func TestOneCommandWithMessage(t *testing.T) {

	// Given:
	content := []byte(`
commands:
	/start:
		message:
		   en_gb: "Hello, this bot can help you to count clicks (CHANGE IT)"
		   ru_ru: "Privet, etot bot mozet delatj raznye fishki (CHANGE IT)"
	`)

	// When:
	rootn, err := _parseYaml(content)

	assert.Nil(t, err)
	assert.NotNil(t, rootn)
}
*/
