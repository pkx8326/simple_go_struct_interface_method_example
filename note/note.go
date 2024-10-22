package note

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"note_title"`
	Content   string    `json:"note_content"`
	CreatedAt time.Time `json:"created_at"`
}

func NoteConstructor(notetitle, notecontent string) (newnote Note) {
	return Note{
		Title:     notetitle,
		Content:   notecontent,
		CreatedAt: time.Now(),
	}
}

func (note Note) Display() {
	fmt.Println("Your note title is ", note.Title)
	fmt.Println("Your note content is ", note.Content)
}

func (note Note) Save() (err error) {
	filename := strings.ReplaceAll(note.Title, " ", "")
	filename = strings.ToLower(filename) + ".json"

	json, err := json.Marshal(note)
	if err != nil {
		err = fmt.Errorf("error saving note to json")
		return
	} else {
		fmt.Println("Saving note to json file...")
		return os.WriteFile(filename, json, 0644)
	}
}
