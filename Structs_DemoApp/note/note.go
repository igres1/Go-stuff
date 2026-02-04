package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"` // this is a struct tag linked to the title field , is basically metadata , we wrote it in the specific format that the json needs it to be understood
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Invali input")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note *Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName)
	json, err := json.Marshal(note) //to be able to write the contents of the struct must be accessible thats why the name of the contents of the struct is written in upperCase
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644) // only read for not owners and returns an error the function

}

func (note *Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", note.Title, note.Content) //automatic dereferencing by compiler xd
}
