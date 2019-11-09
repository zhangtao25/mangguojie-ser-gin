package models

import "fmt"

type Note struct {
	Id   int
	Title string
}


func GetNoteser(title string) ( *Note , error) {
	fmt.Print(title)
	var note Note
	db.Where("title = ?", title).Find(&note)

	return &note, nil
}