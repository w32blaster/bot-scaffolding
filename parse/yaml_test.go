package parse

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestOneCommandWithMessage(t *testing.T) {

	// Given:
	content := []byte(`
commands:
  /start:
    message:
      en_gb: "Hello, this bot can help you to count clicks (CHANGE IT)"
      ru_ru: "Privet, etot bot mozet delatj raznye fishki (CHANGE IT)"`)

	// When:
	rootn, err := _parseYaml(content)

	assert.Nil(t, err)
	assert.NotNil(t, rootn)

	// And:
	assert.Equal(t, 1, len(rootn.Commands))
	assert.Equal(t, "Hello, this bot can help you to count clicks (CHANGE IT)", rootn.Commands["/start"].Message["en_gb"])
	assert.Equal(t, "Privet, etot bot mozet delatj raznye fishki (CHANGE IT)", rootn.Commands["/start"].Message["ru_ru"])
}

func TestFewCommandWithMessage(t *testing.T) {

	// Given:
	content := []byte(`
commands:
  /start:
    message:
      en_gb: "this is start"
      ru_ru: "это начало"
  /help:
    message:
      en_gb: "this is help"
      ru_ru: "это помощь"`)

	// When:
	rootn, err := _parseYaml(content)

	assert.Nil(t, err)
	assert.NotNil(t, rootn)

	// and:
	assert.Equal(t, 2, len(rootn.Commands))
	assert.Equal(t, "this is start", rootn.Commands["/start"].Message["en_gb"])
	assert.Equal(t, "это начало", rootn.Commands["/start"].Message["ru_ru"])

	assert.Equal(t, "this is help", rootn.Commands["/help"].Message["en_gb"])
	assert.Equal(t, "это помощь", rootn.Commands["/help"].Message["ru_ru"])
}
