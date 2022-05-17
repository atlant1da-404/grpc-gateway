package repository

import (
	"fmt"
	"grpc-gateway/internal/dto"
	"grpc-gateway/pkg/gen/notes"
)

type NotesRepository interface {
	GetNotes(id uint64) (*notes.NotesResponse, error)
	CreateNote(in *notes.NotedCreateRequest) (*notes.NotedCreateResponse, error)
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

func (n *notesRepository) CreateNote(in *notes.NotedCreateRequest) (*notes.NotedCreateResponse, error) {

	id := incrementNoteId()

	note := dto.Note{
		Id:    uint(id),
		Name:  in.Name,
		Color: in.Color,
		Count: in.Count,
	}

	bNote := note.Marshal()
	result, err := rdb.Set(ctx, fmt.Sprintf("%s=%d", noteCode, id), bNote, 0).Result()
	if err != nil {
		return nil, err
	}

	return &notes.NotedCreateResponse{Result: result}, nil
}
