package parse

type (

	// Button is one button that can be rendered in a chat
	Button struct {
		Label  map[string]string `yaml:"label"`
		Action string            `yaml:"action"` // the function name to be called when this button is clicked
	}

	// Row is one row of a buttons
	Row struct {
		Buttons []Button `yaml:"row"`
	}

	// Command is a message to a chat starting with "/" symbol. Ex. "/help" or "/start"
	Command struct {
		// Message just one text message that can be send to a chat
		Message    map[string]string `yaml:"message"`
		ButtonRows []Row             `yaml:"buttons"`
	}

	Root struct {
		Commands map[string]Command `yaml:"commands"`
	}
)
