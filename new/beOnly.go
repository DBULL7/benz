package new

import (
	// "fmt"
	// "io"
)


// BeOnly entry point.
func BeOnly() {
	name := GetName()
	database := Database()
	serverTesting := ServerTesting()
	
	if name != "" && database != "" && serverTesting != "" {
		create(name, database, serverTesting)
	} else {
		return 
	}
	// use text template for importing different databases this will increase flexiblity as a whole 
	// will need to go get packages 
}

func create(name, database, testing string) {
	CreateFile("./new/files/server.txt", name + "/server.go")

}

func goGet(projects string) {

}