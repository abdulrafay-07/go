package main

import (
	"fmt"
	"os"
)

func fileLen(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return 0, err
	}

	fileSize := fileInfo.Size()

	return int(fileSize), nil
}

func main() {
	filename := "example.txt"
	length, err := fileLen(filename)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("The file %s has %d bytes\n", filename, length)
	}
}
