package main

import "fmt"

// <- on the right, only writes
func write(name string, hello chan<- string) {
	hello <- name
}

// <- on the left, only reads
func read(data <-chan string) {
	fmt.Println(<-data)
}

func main() {
	hello := make(chan string)
	go write("Hello", hello)
	read(hello)
}
