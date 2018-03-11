package main 

import (
	"os"
	"fmt"	
)

func main() {
	switch command := os.Args[1]; command {
	case "new":
		New()
	default:
		fmt.Println("? Try enzo help to see a list of commands")		
	}	
}