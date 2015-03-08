package main

import (
	"fmt"
	"io/ioutil"
)

// utility fuction to list all the files of directory
// in case of sudirectory it should check for depth level
// and should render the depth level subfolder directories
// what should be the logic

func displayFileInfo(dir string, depth int) {

}

func main() {
	// read directory

	fileinfos, err := ioutil.ReadDir("/home/jittakal/")

	fmt.Println("Eorror: ", err)
	//	fmt.Println(fileinfos)

	if err == nil {
		for _, fileinfo := range fileinfos {
			fmt.Printf("Name: %s IsDir: %t\n", fileinfo.Name(), fileinfo.IsDir())
		}
	}
}
