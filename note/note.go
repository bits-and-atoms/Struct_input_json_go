package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct{
	//capital Note so that others can use this by importing note package
	title string
	content string
	createdAt time.Time
}
func ( note *Note) Display(){
	fmt.Println("note title is: %v\nnote content is: %v\n",note.title,note.content)
}
func New(title,content string) (*Note,error){
	if title == "" || content == ""{
		return &Note{}, errors.New("invalid input")
	}
	return &Note{
		title: title,
		content: content,
		createdAt: time.Now(),
	}, nil
}