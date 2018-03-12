package new

import (
	"fmt"
	// "os"
	"io/ioutil"
	// "path/filepath"

	"path"
	"log"
	"runtime"
)

// CreateFile takes file path of a local file and copies it to the output path 
func CreateFile(fileSrc, output string) {

	  //  _, filename, _, ok := runtime.Caller(0)
    // if !ok {
    //     panic("No caller information")
    // }
		// fmt.Printf("Filename : %q, Dir : %q\n", filename, path.Dir(filename))
		

    // ex, err := os.Executable()
    // if err != nil {
    //     panic(err)
    // }
    // exPath := filepath.Dir(ex)
		// fmt.Println(exPath)
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
			panic("No caller information")
	}
	fmt.Printf("Filename : %q, Dir : %q\n", filename, path.Dir(filename))


	test, err := ioutil.ReadFile(path.Join(filename, "../files/server.text"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(test))
	// fmt.Println(path.Join(filename, "./files"))

}
	
