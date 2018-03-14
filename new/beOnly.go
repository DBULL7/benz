package new


// BeOnly entry point.
func BeOnly(name string) {

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
	Mkdir(name + "/server")
	Mkdir(name + "/server/controllers")
	createServer(name, server)
	// create models folder go get the db type
	// make testing folder go get the testing lib 
	// make gitignore
	// also ask about creating docker file 

}

func createServer(name, server string) {
	switch server {
	case "Chi":
		chi(name)
	case "HTTP Router":
		http(name)
	case "Gin (recommended)":
		gin(name)
	case "Echo":
		echo(name)
	case "Gorilla mux":
		gorilla(name)
	}
}

func chi(name string) {
	CreateFile("../files/chi/.chi.go", name + "/server/server.go")
	CreateFile("../files/chi/.routes.go", name + "/routes.go")

}

func http(name string) {
	CreateFile("../files/http/.http_router.go", name + "/server/server.go")
	CreateFile("../files/http/.routes.go", name + "/server/routes.go")
}

func gin(name string) {
	CreateFile("../files/gin/.gin.go", name + "/server/server.go")
	CreateFile("../files/gin/.routes.go", name + "/server/routes.go")
}

func echo(name string) {
	CreateFile("../files/echo/.echo.go", name + "/server/server.go")
	CreateFile("../files/echo/.routes.go", name + "/server/routes.go")
}

func gorilla(name string) {
	CreateFile("../files/gorilla/.gorilla.go", name + "/server/server.go")
	CreateFile("../files/gorilla/.routes.go", name + "/server/routes.go")
}


