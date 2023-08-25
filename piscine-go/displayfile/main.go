package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := make([]byte, 14)
	_, err = file.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
	file.Close()
}
