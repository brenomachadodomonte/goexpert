package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/* CREATE FILE */
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	/* WRITE FILE */
	//size, err := f.WriteString("Hello, World!")
	size, err := f.Write([]byte("Hello, World!"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("File created. Size: %d bytes\n", size)
	f.Close()

	/* READ FILE */
	file, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file))

	/* READ FILE PER FILE */
	fileLines, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(fileLines)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	/*  REMOVE FILE */
	err = os.Remove("file.txt")
	if err != nil {
		panic(err)
	}
}
