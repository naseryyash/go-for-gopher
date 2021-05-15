package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// os.Args is of type []string, these are the values passed through
	// CLI to our main function - go run main.go myfile.txt
	// myfile.txt is a file in the same folder.
	fileName := os.Args[1]

	// os.Open(filename string) (*File, error) - signature of os.Open()
	file, err := os.Open(fileName)

	// if unable to open the file, return an error, terminate the program
	if err != nil {
		fmt.Println("Error while opening the file:", err)
		os.Exit(1)
	}

	// since file is of type *File, and File type implements both Reader & Writer interface
	// and the signature of io.Copy is like io.Copy(wr Writer, rd Reader)
	io.Copy(os.Stdout, file)
}
