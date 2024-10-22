package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct {
	TodoContent string `json:"todo_content"`
}

func TodoConstructor(todo string) (newtodo Todo) {
	return Todo{
		TodoContent: todo,
	}
}

func (todo Todo) Display() {
	fmt.Println("Your todo content is ", todo.TodoContent)
}

func (todo Todo) Save() (err error) {
	filename := "todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		err = fmt.Errorf("error saving todo to json")
		return
	} else {
		fmt.Println("Saving todo to json file...")
		return os.WriteFile(filename, json, 0644)
	}
}
