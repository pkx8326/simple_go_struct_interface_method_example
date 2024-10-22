package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/struct_interface_02/note"
	"example.com/struct_interface_02/todo"
)

type saver interface {
	Save()
}

type displayer interface {
	Display()
}

func main() {
	notetitle := getinfo("Please input the note's title :")
	notecontent := getinfo("Pleae input the note's content :")
	todocontent := getinfo("Please input the todo :")

	newnote := note.NoteConstructor(notetitle, notecontent)
	newtodo := todo.TodoConstructor(todocontent)

	newnote.Display()
	newtodo.Display()

	err := newnote.Save()
	if err != nil {
		fmt.Println(err)
	}
	err = newtodo.Save()
	if err != nil {
		fmt.Println(err)
	}

}

//////////////////////////////////////////////////////

func getinfo(prompt string) (info string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		info, _ = reader.ReadString('\n')
		info = strings.TrimSpace(info)
		if info == "" {
			fmt.Println("This information cannot be left blank. Please try again.")
			continue
		} else {
			return
		}
	}
}
