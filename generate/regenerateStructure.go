package generate

import (
	"log"
	"os"

	"github.com/w32blaster/bot-scaffolding/parse"

	"github.com/w32blaster/bot-scaffolding/templates"
)

// OverrideStructure regenerate sctructire from the YAML file
func OverrideStructure(rootDir string) {

	// parse YAML file
	root, err := parse.Parse("parse.bot.yaml")
	if err != nil {
		log.Fatal(err)
	}

	commandsDirPath := rootDir + string(os.PathSeparator) + "commands"

	// re-generate commands
	generateFile(commandsDirPath, "commands.generated.go", templates.CommandsFileComntent, root)

}
