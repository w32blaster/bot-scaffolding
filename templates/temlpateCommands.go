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

{{ range $key, $value := .Commands }}
// On{{ commandNameCapital $key }}CommandCalled process the "{{ $key }}" command 
func (m *MyCommandProcessor) On{{ commandNameCapital $key }}CommandCalled(command string) {
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
	{{ range $key, $value := .Commands }} 
	On{{ commandNameCapital $key }}CommandCalled(command string)
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
	{{ range $key, $value := .Commands }}
		case "{{ commandName $key }}":
			process{{ commandNameCapital $key }}Command(command)
		}
	{{ end }}
}

{{ range $key, $value := .Commands }}
// GENERATED
// process the {{ $key }} command
func process{{ commandNameCapital $key }}Command(command string) {

	// .. here comes some logic

	// call custom logic
	impl.On{{ commandNameCapital $key }}CommandCalled(command)

	// ... here we render UI
}

{{ end }}

`
