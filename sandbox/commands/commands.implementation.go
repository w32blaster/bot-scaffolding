
package commands

// Please implement these functions and white here your logic,
// you can edit this file in any way you prefer

// MyCommandProcessor struct that implements processor interface
type MyCommandProcessor struct{}

func New() CommandProcessor {
	return MyCommandProcessor{}
}


// OnHelpCommandCalled process the "/help" command 
func (m *MyCommandProcessor) OnHelpCommandCalled(command string) {
	// your custom logic here
}


// OnStartCommandCalled process the "/start" command 
func (m *MyCommandProcessor) OnStartCommandCalled(command string) {
	// your custom logic here
}



