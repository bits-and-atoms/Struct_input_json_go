package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct{
	//capital Note so that others can use this by importing note package
	Title string
	Content string
	CreatedAt time.Time
}
func ( note *Note) Display(){
	fmt.Println("note title is: ",note.Title,"\nnote content is: ",note.Content,"\n")
}
func (note * Note) Save() error{
	fileName := strings.ReplaceAll(note.Title," ","_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(note)
	//please note that if content of Note are not public (starting with capital)
	//then marshal cant create them in json
	if err != nil{
		return err
	}
	return os.WriteFile(fileName,json,0644)
}
func New(title,content string) (*Note,error){
	if title == "" || content == ""{
		return &Note{}, errors.New("invalid input")
	}
	return &Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil
}