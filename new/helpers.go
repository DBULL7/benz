package new

import (
	// "fmt"
	"os"
	"io/ioutil"
	// "path/filepath"
	// "io"
	"path"
	"log"
	"runtime"
	"strings"
	"github.com/spf13/afero"	
	"github.com/Jeffail/gabs"
	"os/exec"
	"fmt"
)


// CreateFile takes file path of a local file and copies it to the output path 
func CreateFile(fileSrc, output string) {

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
			panic("No caller information")
	}


	file, err := ioutil.ReadFile(path.Join(filename, fileSrc))
	if err != nil {
		log.Fatal(err)
	}


	appFS := afero.NewOsFs()
	error := afero.WriteFile(appFS, output, file, 0644)
	if error != nil {
		log.Fatal(error)
	}
}

//Mkdir helper function to make a directory  
func Mkdir(name string) {
	path := name
	err := os.MkdirAll(path, 0700)

	if err != nil {
		log.Println("Error creating directory")
		log.Println(err)
		os.Exit(1)
	}
}

// GetPath returns users github.com/user/ path
func GetPath(projectPath string) string {
  dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	path := ""
	if strings.Contains(dir, "github.com") {
		path = strings.Split(dir, "src/")[1] + "/" + projectPath
	} else {
		path = projectPath 
	}
	return path
}

// InstallJS npm packages
// requires array of strings. Requires directory to install packages into.
func InstallJS(packages []string, projectDir string) {
	command := exec.Command("yarn", packages...)
	command.Dir = "./" + projectDir
	output, err := command.Output()
	if err != nil {
			log.Println(err)
	}
	fmt.Printf("%s", output)
}

// InstallE2E installs end to end testing for applications. 
// Takes the string e2e for determining which testing system to use.
// Takes string name which is the name of the project for creating files.
func InstallE2E(e2e, name string) {
	if e2e == "Cypress" {
		cypress(name)
	} else if e2e == "Test Cafe" {
		testcafe(name)
	}
}

func cypress(name string) {
	Mkdir(name + "/cypress")
	Mkdir(name + "/cypress/integration")
	CreateFile("../files/test/e2e/cypress.js", name + "/cypress/integration/test.js")
	packages := []string{"add", "cypress", "--dev"}
	InstallJS(packages, name)
	AddScriptToPackageJSON("e2e", "cypress open", name)
}

func testcafe(name string) {
	Mkdir(name + "/test")
	Mkdir(name + "/test/e2e")
	CreateFile("../files/test/e2e/testcafe.js", name + "/test/e2e/test.js")
	packages := []string{"add", "testcafe", "--dev"}
	InstallJS(packages, name)
	AddScriptToPackageJSON("e2e", "testcafe chrome test/e2e", name)
}

// AddScriptToPackageJSON Creates a new script in a package.json file.
// Accepts the script key (ie "test") with it's corresponding value and the project name.
func AddScriptToPackageJSON(scriptKey, scriptValue, projectName string) {
	json, _ := gabs.ParseJSONFile(projectName + "/package.json")
	json.SetP(scriptValue, "scripts."+ scriptKey)

	jsonOutput := json.StringIndent("", "  ")
	appFS := afero.NewOsFs()
	error := afero.WriteFile(appFS, projectName + "/package.json", []byte(jsonOutput), 0644)
	if error != nil {
		fmt.Println(error)
	}	
}

