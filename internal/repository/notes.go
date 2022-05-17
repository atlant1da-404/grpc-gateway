package repository

import (
	"fmt"
	"grpc-gateway/internal/dto"
	"grpc-gateway/pkg/gen/notes"
)

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

	result, err := rdb.Get(ctx, fmt.Sprintf("%s=%d", noteCode, id)).Result()
	if err != nil {
		return nil, err
	}

	note := dto.Note{}
	note.Unmarshal([]byte(result))

	return &notes.NotesResponse{
		Name:  note.Name,
		Color: note.Color,
		Count: note.Count,
	}, nil
}
