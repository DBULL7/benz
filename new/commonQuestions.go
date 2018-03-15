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
	prompt := &survey.Select {
			Message: "Choose a Testing Framework:",
			Options: []string{"Go Standard Library", "Testify", "Go Convey", "None"},
	}
	survey.AskOne(prompt, &serverTesting, nil)
	return serverTesting
}

// BackendServer choose type of backend server 
func BackendServer() string {
	server := ""
	prompt := &survey.Select {
			Message: "Choose a Server: ",
			Options: []string{ "Gin (recommended)", "Echo", "Chi", "Gorilla mux", "HTTP Router" },
	}
	survey.AskOne(prompt, &server, nil)
	return server
}

// ReactTesting choose if you want to test your react files
func ReactTesting() bool {
	testing := false
	prompt := &survey.Confirm {
		Message: "Do you want to use Enzyme and Jest to test your React components",
	}
	survey.AskOne(prompt, &testing, nil)
	return testing 
}

// E2e select a Javascript e2e testing framework
func E2e() string {
	tool := ""
	prompt := &survey.Select {
		Message: "e2e Javascript Testing Framework",
		Options: []string{ "Cypress", "Test Cafe", "None" },
	}
	survey.AskOne(prompt, &tool, nil)
	return tool
}

// Backend asks user if they need backend 
func Backend() bool {
	backend := false 
	prompt := &survey.Confirm {
		Message: "Do need a backend",
	}
	survey.AskOne(prompt, &backend, nil)
	return backend
}