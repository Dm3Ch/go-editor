package main

import (
	"fmt"
	"os"

	editor "github.com/Dm3Ch/go-editor"
)

func main() {
	err := editor.RunEditor("./test.txt")
	if err != nil {
		fmt.Println("Error while editing file")
		fmt.Println(err)
		os.Exit(1)
	}
}
