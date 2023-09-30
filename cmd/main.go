package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	directory := "./"

	files, err := os.Open(directory)
	if err != nil {
		fmt.Println("error opening directory:", err)
		return
	}
	defer files.Close()

	fileInfos, err := files.Readdir(-1)
	if err != nil {
		fmt.Println("error reading directory:", err)
		return
	}

	fmt.Printf("Converting files")
	fmt.Println()

	fileCount := len(fileInfos)

	for i, fileInfos := range fileInfos {
		fmt.Printf("Processing file %v of %v", i+1, fileCount)
		fmt.Println()

		fileParts := strings.Split(fileInfos.Name(), ".")
		if fileParts[1] != "HEIC" {
			continue
		}
		cmd := exec.Command("magick", fileInfos.Name(), fileParts[0]+".png")
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("All done!")
}
