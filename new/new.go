package new 

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
			Options: []string{ "SPA" , "Multi Page Applications" , "MVC" , "API Backend" },
		},
	},
}

func typeOfProject() {
	answer := struct {
			Type          string                  
	}{}

	err := survey.Ask(qs, &answer)
	if err != nil {
			fmt.Println(err.Error())
			return
	}
	name := GetName()
	switch answer.Type {
	case "SPA":
		spa()
	case "Multi Page Applications":
		multi()
	case "MVC":
		MVC()
	case "API Backend":
		BeOnly(name)
	default: 
		fmt.Println("No option selected?")
	}
}


func spa() {
	fe := chooseFE()
	fmt.Println(fe)
	if fe == "React" {
		ReactSPA()
	} else {
		VueSPA()
	}
}

func multi() {
	fe := chooseFE()
	fmt.Println(fe)
}

func chooseFE() string {
	fe  := ""
	prompt := &survey.Select {
			Message: "Use React or Vue:",
			Options: []string{ "React", "Vue" },
	}
	survey.AskOne(prompt, &fe, nil)
	return fe
} 
