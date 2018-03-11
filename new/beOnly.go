package new

import (
	"fmt"
)

// BeOnly entry point.
func BeOnly() {
	name := GetName()
	database := Database()
	serverTesting := ServerTesting()
	
	if name != "" && database != "" && serverTesting != "" {
		// run func
		fmt.Println("I fired") 
	} else {
		return 
	}
	// use text template for importing different databases this will increase flexiblity as a whole 
	// will need to go get packages 
}