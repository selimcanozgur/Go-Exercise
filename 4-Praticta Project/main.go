package main

import (
	"bufio"
	"fmt"
	"os"

	"example.com/note/note"
)


func main() {
	title, content := getNoteData() 
	
	userNote, err := note.New(title,content)

	
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()


}

func getNoteData () (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString()
	return value
}