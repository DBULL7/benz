package new

import (
	// "fmt"
	// "io"
	// "log"
	// "os"
)


// BeOnly entry point.
func BeOnly() {
	name := GetName()
	server := BackendServer()
	database := Database()
	serverTesting := ServerTesting()
	

	if name != "" && server != "" && database != "" && serverTesting != "" {
		create(name, server, database, serverTesting)
	} else {
		return 
	}
}

func create(name, server, database, testing string) {
	Mkdir(name)
	createServer(name, server)
}

func createServer(name, server string) {
	switch server {
	case "HTTP Router":
		CreateFile("../files/", name + "/server.go")
	case "Gin":
		CreateFile("../files/.gin.go", name + "/server.go")
	case "Echo":
		CreateFile("../files/.echo.go", name + "/server.go")
	case "Gorilla mux":
		CreateFile("../files/.gorilla.go", name + "/server.go")
	case "Go standard library":
		CreateFile("../files/.go_server.go", name + "/server.go")
	case "Fast Http":
		CreateFile("../files/.fast_http.go", name + "/server.go")
	}
}

func goGet(projects string) {

}