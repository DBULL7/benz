package new  

import (
	"fmt"
	// "os"
	// "strings"
	// "io/ioutil"
	"github.com/spf13/afero"
	"github.com/Jeffail/gabs"
)



// MVC project creation entry point.
func MVC() {
	test, _ := gabs.ParseJSONFile("package.json")
	// fmt.Println(test)
	test.SetP("hell yeah","scripts.test")
	// value := test.Search("scripts")

	// fmt.Println(value)
	jsonOutput := test.StringIndent("", "  ")
	appFS := afero.NewOsFs()
	error := afero.WriteFile(appFS, "package.json", []byte(jsonOutput), 0644)
	if error != nil {
		fmt.Println(error)
	}	
}