package main

import (
	"fmt"
	"ubyjerome/dystra/routines"
	"ubyjerome/dystra/utilities"
)

func main() {
	fmt.Println("Project Entry Point")
	routine.ScanForIncludedFolders()

	file, err := utilities.FetchFile("./.dystra.config.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(file)
}
