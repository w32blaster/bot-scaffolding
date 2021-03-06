package templates

// SimpleInitialYAMLFile very simple YAML file for the initial project that will be generated
// on a new project creation
const SimpleInitialYAMLFile = `
# 
# This is the sample code for simple bot that counts clicks
#
commands:
	- /start:
		- message:
			- en-gb: "Hello, this bot can help you to count clicks (CHANGE IT)"
			- ru-ru: "Privet, etot bot mozet delatj raznye fishki (CHANGE IT)"
	- /help:
		- message:
			- en-gb: "/start - start the bot from the beginning\n
			          /help - show the help"
			- ru-ru: "/start - na4atj bot\n
					  /help - pokazatj help"
	- /count:
		- message:
		    - en-gb: "Clicks count is 0"
			- ru-ru: "Kol-vo klikov 0"
		- buttons:
			- row:
			 	- button:
					action: increaseCount
					label: 
						- en-gb: "Plus"
						- ru-ru: "Plus"
				- button:
					action: decreaseCount
					label: 
						- en-gb: "Minus"
						- ru-ru: "Minus" 

`
