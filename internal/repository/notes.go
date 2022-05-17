package repository

import "grpc-gateway/pkg/gen/notes"

type NotesRepository interface {
	GetNotes(id uint64) (*notes.NotesResponse, error)
}

type notesRepository struct {
	NotesRepository
}

func NewNotesRepository() *notesRepository {
	return &notesRepository{}
}

func (n *notesRepository) GetNotes(id uint64) (*notes.NotesResponse, error) {

	return &notes.NotesResponse{
		Name:  "andrik",
		Color: "black",
		Count: 53,
	}, nil
}
