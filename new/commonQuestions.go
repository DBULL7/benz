package new

import (
	"gopkg.in/AlecAivazis/survey.v1"
)

// GetName asks for and returns the new project name. 
func GetName() string {
	name := ""
	prompt := &survey.Input{
			Message: "Project name:",
	}
	survey.AskOne(prompt, &name, nil)
	return name 
}

// Database asks user for what type of DB if any they want to use.
func Database() string {
	db := ""
	prompt := &survey.Select{
			Message: "Choose a Database:",
			Options: []string{"Postgres", "MongoDB", "None"},
	}
	survey.AskOne(prompt, &db, nil)
	return db 
}

// ServerTesting returns user selection of server testing tools 
func ServerTesting() string {
	serverTesting := ""
	prompt := &survey.Select{
			Message: "Choose a Testing Framework:",
			Options: []string{"Go Standard Library", "Testify", "Go Convey", "None"},
	}
	survey.AskOne(prompt, &serverTesting, nil)
	return serverTesting
}