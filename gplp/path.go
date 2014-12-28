package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {
	// declare array
	componenets := []string{"a", "path",
		".", "with", "relative", "elements"}

	// join path
	// variadic function
	// The ... suffix on the slice indicates that
	// it should be passed as the variadic argument set,
	// rather than as a single argument
	path := path.Join(componenets...)

	// print path
	fmt.Printf("Path: %s\n", path)

	// split file list
	decomposed := splitPath(path)

	// iterate the file list and print it
	for i, dir := range decomposed {
		fmt.Println(i)
		fmt.Printf("%s%c", dir, filepath.Separator)
		i += 1
	}

	fmt.Printf("\n")
}

func splitPath(path string) []string {
	return filepath.SplitList(path)
}
