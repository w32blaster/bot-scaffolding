
package commands
// GENERATED

// CommandProcessor interface for the command processor.
// GENERATED 
// Should be implemented by a custom logic, because UI widget rendering will be done by generated code from YAML file
type CommandProcessor interface {
	
	ProcessHelpCommand(command string)
	
	ProcessStartCommand(command string)
	
}

var impl CommandProcessor

func SetImpl(i CommandProcessor) {
	impl = i
}

// ProcessCommand process one command. A command must start with a slash and may contain
// some arguments (such as "/help" or "/dosmth arg1 arg2")
func ProcessCommand(command string) {

	switch command {
			
		case "/Help":
			processHelpCommand(command)
		}
			
		case "/Start":
			processStartCommand(command)
		}
	
}


// GENERATED
// process the Help command
func processHelpCommand(command string) {

	// .. here comes some logic

	// call custom logic
	impl.ProcessHelpCommand(command)

	// ... here we render UI
}


// GENERATED
// process the Start command
func processStartCommand(command string) {

	// .. here comes some logic

	// call custom logic
	impl.ProcessStartCommand(command)

	// ... here we render UI
}



