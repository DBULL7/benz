package main 

import (
	"fmt"
	"gopkg.in/AlecAivazis/survey.v1"
)

// New is the key function that handles creating new projects
func New() {
	typeOfProject()
}

var qs = []*survey.Question {
	{
		Name: "Type",
		Prompt: &survey.Select{
			Message: "Choose a Project Type:",
			Options: []string{"SPA","Multi Page Applications","MVC", "Backend Only"},
		},
	},
}

func typeOfProject() {
	answers := struct {
			Type          string                  
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
			fmt.Println(err.Error())
			return
	}
	fmt.Println(answers.Type)
}