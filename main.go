package main

import (
	"flag"

	"github.com/w32blaster/bot-scaffolding/generate"
)

func main() {

	// the first argument is expected to specify what exactly we want to generate.
	pOperation := flag.String("o", "generate", "What exatly we want to generate: the initial project structure or update the bot map; default is 'generate' ")
	pTargetDir := flag.String("t", ".", "Target folder where the code should be placed; default is current directory")
	flag.Parse()

	switch *pOperation {
	case "init":
		{

			// generate the initial content. This command should be called only once to create a new project
			generate.InitialStructure(*pTargetDir)
		}
	case "generate":
		{

			// regenerate the content. Should be called
			generate.OverrideStructure(*pTargetDir)
		}
	}

}
