package note_service

import (
	"github.com/zhangtao25/mangostreet-ser-gin/models"
)

type Note struct {
	Id   int
	Title string
}

func (a *Note) Get() (*models.Note, error) {
	note, err := models.GetNoteser(a.Title)
	if err != nil {
		return nil, err
	}

	return note, err
}