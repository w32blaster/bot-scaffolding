package templates

// CommandsInitCustomImplFileComntent the initial file for implementation.
// Should be called only one on the project generation to help a developer to start
// writing implementations
const CommandsInitCustomImplFileComntent = `
package commands

// Please implement these functions and white here your logic,
// you can edit this file in any way you prefer

// MyCommandProcessor struct that implements processor interface
type MyCommandProcessor struct{}

func New() CommandProcessor {
	return MyCommandProcessor{}
}

{{ range .commands }}
// Process{{ . }}Command process the "/{{ . }}" command 
func (m *MyCommandProcessor) Process{{ . }}Command(command string) {
	// your custom logic here
}

{{ end }}

`

// CommandsFileComntent generates the processors for the commands
const CommandsFileComntent = `
package commands
// GENERATED

// CommandProcessor interface for the command processor.
// GENERATED 
// Should be implemented by a custom logic, because UI widget rendering will be done by generated code from YAML file
type CommandProcessor interface {
	{{ range .commands }}
	Process{{ . }}Command(command string)
	{{ end }}
}

var impl CommandProcessor

func SetImpl(i CommandProcessor) {
	impl = i
}

// ProcessCommand process one command. A command must start with a slash and may contain
// some arguments (such as "/help" or "/dosmth arg1 arg2")
func ProcessCommand(command string) {

	switch command {
	{{ range .commands }}		
		case "/{{ . }}":
			process{{ . }}Command(command)
		}
	{{ end }}
}

{{ range .commands }}
// GENERATED
// process the {{ . }} command
func process{{ . }}Command(command string) {

	// .. here comes some logic

	// call custom logic
	impl.Process{{ . }}Command(command)

	// ... here we render UI
}

{{ end }}

`
