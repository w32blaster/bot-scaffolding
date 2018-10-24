package generate

import (
	"html/template"
	"log"
	"os"

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

	commandsTemplateData := map[string]interface{}{
		"commands": []string{"Help", "Start"},
	}

	generateFile(commandsDirPath, "commands.generated.go", templates.CommandsFileComntent, commandsTemplateData)
	generateFile(commandsDirPath, "commands.implementation.go", templates.CommandsInitCustomImplFileComntent, commandsTemplateData)

}

func generateFile(rootDir, resultFileName, templateBody string, data interface{}) error {
	f, err := os.Create(rootDir + string(os.PathSeparator) + resultFileName)
	if err != nil {
		log.Fatal(err)
	}

	t := template.Must(template.New(resultFileName).Parse(templateBody))

	err = t.Execute(f, data)
	if err != nil {
		log.Fatalf("error: %v", err)
		return err
	}

	return nil
}
