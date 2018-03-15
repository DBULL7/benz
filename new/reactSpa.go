package new

import (
	// "fmt"
)

// ReactSPA entry point for building create-react-app style applications
func ReactSPA(name string) {
	reactTesting := ReactTesting()
	e2e := E2e()
	if backend := Backend(); backend == true {
		BeOnly(name, "html")
	} else {
		// create html file in root of project. 
		// package json needs to have command for webpack dev server with hot reload 
	}

	reactTestingSetup(reactTesting)
	installE2E(e2e)
	Mkdir(name + "/src")
	// need to npm install everything 
	// .babelrc
	// .gitignore
	// readme 
	// webpack.js
	// src
	// src/Home
	// src/Home/Home.js
	// dist 
}

func reactTestingSetup(confirmed bool) {
	if (confirmed) {
		// create specific package.json file for jest 
		// 
	} else {
		// normal package.json 
	}
}

func installE2E(e2e string) {
	if e2e == "Cypress" {

	} else if e2e == "Test Cafe" {

	}
}