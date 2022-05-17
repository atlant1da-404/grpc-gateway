package service

import (
	"errors"
	"grpc-gateway/internal/repository"
	"grpc-gateway/pkg/gen/notes"
)

type NotesService interface {
	GetNotes(id uint64) (*notes.NotesResponse, error)
	CreateNote(in *notes.NotedCreateRequest) (*notes.NotedCreateResponse, error)
	UpdateNote(in *notes.NoteUpdateRequest) (*notes.NoteUpdateResponse, error)
	DeleteNote(in *notes.NoteDeleteRequest) (*notes.NoteDeleteResponse, error)
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

func (n *notesService) CreateNote(in *notes.NotedCreateRequest) (*notes.NotedCreateResponse, error) {

	if len([]rune(in.Name)) <= 0 {
		return nil, errors.New("not valid name")
	}

	if len([]rune(in.Color)) <= 0 {
		return nil, errors.New("not valid color")
	}

	result, err := n.notesRepository.CreateNote(in)
	return result, err
}

func (n *notesService) UpdateNote(in *notes.NoteUpdateRequest) (*notes.NoteUpdateResponse, error) {

	if len([]rune(in.Name)) <= 0 {
		return nil, errors.New("not valid name")
	}

	if len([]rune(in.Color)) <= 0 {
		return nil, errors.New("not valid color")
	}

	result, err := n.notesRepository.UpdateNote(in)
	return result, err
}

func (n *notesService) DeleteNote(in *notes.NoteDeleteRequest) (*notes.NoteDeleteResponse, error) {

	result, err := n.notesRepository.DeleteNote(in)
	return result, err
}
