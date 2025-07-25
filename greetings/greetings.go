package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func Print(message string) string {
	return message
}

func RunLoop(number int) {
	for number < 10 {
		number = number + 1
	}
	fmt.Println(number)
}
