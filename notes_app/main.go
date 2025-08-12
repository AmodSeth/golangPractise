package main

import (
	"bufio"

	"fmt"
	"os"
	"strconv"
	"strings"

	"example.com/notes/note"
	"example.com/notes/todo"
)



//a contract that a certain value has that method for sure
type saver interface{
	Save() error
}

func getNoteData() (string ,string,int)  {
	//title 
	title:= getUserInput("Notes Title: ")
	//content
	content := getUserInput("Notes Content: ")

	priority_string := getUserInput("Notes Priority")
	
	priority :=stringtoint(priority_string)


	return title,content,priority

}


func getUserInput(prompt string) (string){
	fmt.Println(prompt);

	//only for single word
	// var input string
	// fmt.Scanln(&input)	
	// return input
	
	//***** new line inputs  ************//

	//input from command line
	reader:= bufio.NewReader(os.Stdin)
	// 
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}
	//now remove the delimeter

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text




}



func stringtoint(priority string) (int){
	priority_int , err := strconv.Atoi(priority)
	if err != nil {
		return 0 
	}
	return priority_int
}

func main() {
	title , content ,priority := getNoteData();

	//get the todo
	new_todo , err := todo.NewTodo(getUserInput("Please enter the todo Text"))

	if err != nil {
		return 
	}
	new_todo.DisplayNewNote()
	
	if new_todo.Save() != nil {
		fmt.Println("error", err)
	}else{
		fmt.Println("Todo saved successfuly")
	}
	
	user_note , err := note.NewNote(title,content,priority)
	if err != nil {
		fmt.Println("Error",err)
		return
	}

	user_note.DisplayNewNote()



	//savin the note

	if user_note.Save() != nil {
		fmt.Println("error", err)
	}else{
		fmt.Println("File save successfuly")
	}

}