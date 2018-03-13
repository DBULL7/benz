package main 

import (

)

func main() {
  app := InitializeRoutes()
	
	app.Run(":3000")
}

