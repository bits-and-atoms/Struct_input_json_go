package structinputjsongo

import (
	"errors"
	"fmt"

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
}

func getNoteData() (string,string){
	title := getUserInput("Note title-")
	content := getUserInput("Note content-")
	return title,content
}
func getUserInput(prompt string) (string){
	fmt.Println(prompt)
	var val string
	fmt.Scanln(&val)
	return val

}
