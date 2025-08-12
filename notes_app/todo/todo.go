package todo

import (
	"encoding/json"

	"fmt"
	"os"
)

type Todo struct{
	Text string `json: "text"`
}

//contructor
func NewTodo(content string ) (Todo , error) {
	
	if content == "" {
		return Todo{}, fmt.Errorf("content cannot be Empty")
	}

	return Todo{
	
		Text: content,
	
	}, nil


}

//display method

func (todo Todo) DisplayNewNote() {
    fmt.Printf(`
	
	Content: %v
	
	
	`, todo.Text)
}

func (todo Todo) Save() error {
	//replace white spaces with _
	filename := "todo.json"

	
	
	//creating a directory
	json , err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(filename, json, 0644)

}