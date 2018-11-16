package generate

import (
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/w32blaster/bot-scaffolding/parse"
	"github.com/w32blaster/bot-scaffolding/templates"
)

// InitialStructure generate inital structure for a new project
func InitialStructure(rootDir string) {

	var emptpyData interface{}

	// 1. Generate sample yaml file
	generateFile(rootDir, "bot.yaml", templates.SimpleInitialYAMLFile, emptpyData)

	// 2. Generate main.go file
	generateFile(rootDir, "main.go", templates.MainFile, emptpyData)

	// 3. Generate sample processing file and implementations (with folders)
	commandsDirPath := rootDir + string(os.PathSeparator) + "commands"
	os.Mkdir(commandsDirPath, 0755)

	/*
		commandsTemplateData := map[string]interface{}{
			"commands": []string{"Help", "Start"},
		}
	*/
	rootDefaultCommands := parse.Root{
		Commands: map[string]parse.Command{
			"/start": parse.Command{
				Message: map[string]string{
					"en_gb": "test",
				},
			},
			"/help": parse.Command{
				Message: map[string]string{
					"en_gb": "help",
				},
			},
		},
	}

	generateFile(commandsDirPath, "commands.generated.go", templates.CommandsFileComntent, rootDefaultCommands)
	generateFile(commandsDirPath, "commands.implementation.go", templates.CommandsInitCustomImplFileComntent, rootDefaultCommands)

}

func generateFile(rootDir, resultFileName, templateBody string, data interface{}) error {

	f, err := os.Create(rootDir + string(os.PathSeparator) + resultFileName)
	if err != nil {
		log.Fatal(err)
	}

	templateCustomFunctions := template.FuncMap{

		"commandName": func(command string) string {
			return command[1:]
		},

		"commandNameCapital": func(command string) string {
			return strings.Title(command[1:])
		},
	}

	t := template.Must(template.New(resultFileName).Funcs(templateCustomFunctions).Parse(templateBody))

	err = t.Execute(f, data)
	if err != nil {
		log.Fatalf("error: %v", err)
		return err
	}

	return nil
}
