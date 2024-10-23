/*
	Adding a Method & Handling User Input
	- fmt.Scanln() -> does not work with long inputs
	- we need to use bufio.NewReader(os.Stdin)

	1. main.go
	2. note/note.go -> Display()


*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin) // 1.
	text, err := reader.ReadString('\n') 

	// 2.
	if err != nil {
		return ""
	}

	// 3.
	text = strings.TrimSuffix(text, "\n") // trim trailing \n
	text = strings.TrimSuffix(text, "\r") // \r\n 

	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}


func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display() // #

}
