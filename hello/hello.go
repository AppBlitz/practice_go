package main

import "fmt"

func main() {
	if v := 10; v > 11 {
		fmt.Println("The number ", v)
	} else if v < 11 {
		fmt.Println(v)
	}
}
