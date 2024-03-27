package main

import "log"

func panic1() {
	panic("panic1")
}

func panic2() {
	panic("panic2")
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			if r == "panic1" {
				log.Println("panic1 recovered", r)
			}

			if r == "panic2" {
				log.Println("panic2 recovered", r)
			}
		}
	}()

	panic2()
}
