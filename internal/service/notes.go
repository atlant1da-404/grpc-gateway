package service

import (
	"grpc-gateway/internal/repository"
	"grpc-gateway/pkg/gen/notes"
)

type NotesService interface {
	GetNotes(id uint64) (*notes.NotesResponse, error)
}

type notesService struct {
	NotesService
	notesRepository repository.NotesRepository
}

func NewNotesService() *notesService {
	return &notesService{
		notesRepository: repository.NewNotesRepository(),
	}
}

func (n *notesService) GetNotes(id uint64) (*notes.NotesResponse, error) {

	result, err := n.notesRepository.GetNotes(id)
	return result, err
}
