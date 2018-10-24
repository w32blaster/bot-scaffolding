
package commands

// Please implement these functions and white here your logic,
// you can edit this file in any way you prefer

// MyCommandProcessor struct that implements processor interface
type MyCommandProcessor struct{}

func New() CommandProcessor {
	return MyCommandProcessor{}
}


// ProcessHelpCommand process the "/Help" command 
func (m *MyCommandProcessor) ProcessHelpCommand(command string) {
	// your custom logic here
}


// ProcessStartCommand process the "/Start" command 
func (m *MyCommandProcessor) ProcessStartCommand(command string) {
	// your custom logic here
}



