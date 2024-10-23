/*
	Getting User Input



*/

package main

import (
	"errors"
	"fmt"
)

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)

	if value == "" {
		return "", errors.New("invalid input")
	}

	return value, nil
}

// 2. return title, content, error
func getNoteData() (string, string, error) {
	title, err := getUserInput("Note title:") // a. return error

	// b.
	if err != nil {
		return "", "", err
	}

	content, err := getUserInput("Note content:")

	if err != nil {
		return "", "", err
	}

	return title, content, nil
}

func main() {
	title, content, err := getNoteData() // 3.

	// 4.
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(title, content)
}


