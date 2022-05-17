package service

import "grpc-gateway/pkg/gen/notes"

type NotesService interface {
	GetNotes(id uint64) (*notes.NotesResponse, error)
}

type notesService struct {
	NotesService
}

func NewNotesService() *notesService {
	return &notesService{}
}

func (n *notesService) GetNotes(id uint64) (*notes.NotesResponse, error) {

	return &notes.NotesResponse{
		Name:  "andrik",
		Color: "black",
		Count: 53,
	}, nil
}
