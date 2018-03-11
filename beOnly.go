package main

import (
	"fmt"
	// "gopkg.in/AlecAivazis/survey.v1"
)

// BeOnly entry point.
func BeOnly() {
	name := GetName()
	database := Database()
	serverTesting := ServerTesting()
	fmt.Println(name, database, serverTesting)
	

	// use text template for importing different databases this will increase flexiblity as a whole 
	// will need to go get packages 
}