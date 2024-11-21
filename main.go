package main

import "fmt"

func main() {
	counter := 0
	for {
		counter++
		fmt.Printf("%d -- Hello, World!\n", counter)

		if counter == 10 {
			break
		}

	}
}
