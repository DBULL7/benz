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
	// create models folder go get the db type
	// make testing folder go get the testing lib 
}

func createServer(name, server string) {
	switch server {
	case "HTTP Router":
		http(name)
	case "Gin":
		gin(name)
	case "Echo":
		echo(name)
	case "Gorilla mux":
		gorilla(name)
	case "Go standard library":
		goStnd(name)
	case "Fast Http":
		fasthttp(name)
	}
}

func http(name string) {
	CreateFile("../files/http/.http_router.go", name + "/server.go")
	CreateFile("../files/http/.routes.go", name + "/routes.go")
}

func gin(name string) {
	CreateFile("../files/gin/.gin.go", name + "/server.go")
	CreateFile("../files/gin/.routes.go", name + "/routes.go")
}

func echo(name string) {
	CreateFile("../files/echo/.echo.go", name + "/server.go")
	CreateFile("../files/echo/.routes.go", name + "/routes.go")
}

func gorilla(name string) {
	CreateFile("../files/gorilla/.gorilla.go", name + "/server.go")
	CreateFile("../files/gorilla/.routes.go", name + "/routes.go")
}

func goStnd(name string) {
	CreateFile("../files/goStnd/.go_server.go", name + "/server.go")
	CreateFile("../files/goStnd/.routes.go", name + "/routes.go")
}

func fasthttp(name string) {
	CreateFile("../files/fast_http/.fast_http.go", name + "/server.go")
	CreateFile("../files/fast_http/.routes.go", name + "/routes.go")
}