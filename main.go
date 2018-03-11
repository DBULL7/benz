package main 

import (
	"os"
	"fmt"
	"github.com/DBULL7/benz/new"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("? No argument supplied. Try enzo help to see a list of commands.")
		return 
	}
	switch command := os.Args[1]; command {
	case "new":
		new.New()
	default:
		fmt.Println("? Try enzo help to see a list of commands")		
	}	
}