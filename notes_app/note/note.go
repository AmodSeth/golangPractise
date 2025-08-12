package note

import (
	"encoding/json"

	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct{
	Title string
	Content string
	Priority int
	CreatedAt time.Time
}

//contructor
func NewNote(title string , content string , priority int) (Note , error) {
	
	if title == "" || content == "" {
		return Note{}, fmt.Errorf("Tile and Content cannot be Empty")
	}

	return Note{
		Title: title,
		Content: content,
		Priority: priority,
		CreatedAt: time.Now(),
	}, nil


}

//display method

func (n Note) DisplayNewNote() {
    fmt.Printf(`
	Title: %v,
	Content: %v
	priority: %v
	createdAt: %v
	`,n.Title, n.Content , n.Priority , n.CreatedAt)
}

func (n Note) Save() error {
	//replace white spaces with _
	filename := strings.ReplaceAll(n.Title," ", "_")
	fmt.Println("this is getting written before to lower ===> ", filename )
	filename =  strings.ToLower(filename) + ".json"

	fmt.Println("this is getting written after to lower===> ", filename )
	
	//creating a directory
	json , err := json.Marshal(n)

	if err != nil {
		return err
	}

	return os.WriteFile(filename, json, 0644)

}