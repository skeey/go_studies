package main

import (
	"fmt"
	"os"
)

func main() {
	tmp := os.TempDir()
	
	makeFiles(tmp)
	makeFiles(tmp, "test1")
	makeFiles(tmp, "test2", "test3", "test4")
}

func makeFiles(dirBase string, files ...string) {
	for _, name := range files {
		path := fmt.Sprintf("%s/%s.%s", dirBase, name, "txt")
		
		file, err := os.Create(path)

		defer file.Close()

		if err != nil {
			fmt.Printf("Error when trying to create the file %s: %v\n", name, err)
			os.Exit(1)
		}

		fmt.Printf("File %s created.\n", file.Name())
	}
}
