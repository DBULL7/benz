package new

import (
	"fmt"
	// "os"
	"io/ioutil"
	// "path/filepath"
	// "io"
	"path"
	"log"
	"runtime"
	"github.com/spf13/afero"	
)


// CreateFile takes file path of a local file and copies it to the output path 
func CreateFile(fileSrc, output string) {

	
 
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
			panic("No caller information")
	}
	// fmt.Printf("Filename : %q, Dir : %q\n", filename, path.Dir(filename))


	file, err := ioutil.ReadFile(path.Join(filename, fileSrc))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(file))

  // from, err := os.Open(path.Join(filename, fileSrc))
  // if err != nil {
  //   log.Fatal(err)
  // }
  // defer from.Close()

  // to, err := os.OpenFile(path.Join(filename, fileSrc), os.O_RDWR|os.O_CREATE, 0700)
  // if err != nil {
  //   log.Fatal(err)
  // }
	// defer to.Close(
	// fmt.Println(from)
	fmt.Printf("This is the output: %v", output)
	appFS := afero.NewOsFs()
	error := afero.WriteFile(appFS, output, file, 0644)
	if error != nil {
		fmt.Println(error)
	}
	// fmt.Println(to)
  // _, err = io.Copy(to, from)
  // if err != nil {
  //   log.Fatal(err)
  // }

}
	
