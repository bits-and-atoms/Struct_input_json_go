package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bits-and-atoms/Struct_input_json_go/note"
)

func main(){
	title, content := getNoteData()
	userNote, err:= note.New(title,content)
	if err != nil{
		fmt.Println(err)
		return
	}
	userNote.Display()
	err = userNote.Save()
	if err != nil{
		fmt.Println(err)
		return
	}
}

func getNoteData() (string,string){
	title := getUserInput("Note title-")
	content := getUserInput("Note content-")
	return title,content
}
func getUserInput(prompt string) (string){
	fmt.Println(prompt)
	// var val string
	// fmt.Scanln(&val)
	// fmt.Scanln reads input until the first space or newline, 
	// so it only captures a single word
	reader := bufio.NewReader(os.Stdin)
	text , err := reader.ReadString('\n')
	// read till you get a line break also here ' is used not "
	if err != nil{
		return ""
	}
	text = strings.TrimSuffix(text,"\n")
	text = strings.TrimSuffix(text,"\r")
	//from input we want to remove the \n but in windows its not just
	// \n its /r/n so remove both

	return text

}
