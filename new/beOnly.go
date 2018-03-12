package new

import (
	"fmt"
	// "io"
	"log"
	"os"
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
	path := name
	fmt.Println(path)
	err := os.MkdirAll(path, 0700)

	if err != nil {
		log.Println("Error creating directory")
		log.Println(err)
		return
	}
	CreateFile("../files/server.text", name+"/server.go")

}

func goGet(projects string) {

}