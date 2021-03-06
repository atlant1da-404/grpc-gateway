package handler

import (
	"context"
	"grpc-gateway/internal/service"
	"grpc-gateway/pkg/gen/notes"
)

type Notes struct {
	notes.UnimplementedNotesServer
	notesService service.NotesService
}

func NewNotes() *Notes {
	return &Notes{
		notesService: service.NewNotesService(),
	}
}

func (s *Notes) GetNote(ctx context.Context, in *notes.NotesRequest) (*notes.NotesResponse, error) {

	result, err := s.notesService.GetNotes(in.Id)
	return result, err
}

func (s *Notes) CreateNote(ctx context.Context, in *notes.NotedCreateRequest) (*notes.NotedCreateResponse, error) {

	result, err := s.notesService.CreateNote(in)
	return result, err
}

func (s *Notes) UpdateNote(ctx context.Context, in *notes.NoteUpdateRequest) (*notes.NoteUpdateResponse, error) {

	result, err := s.notesService.UpdateNote(in)
	return result, err
}

func (s *Notes) DeleteNote(ctx context.Context, in *notes.NoteDeleteRequest) (*notes.NoteDeleteResponse, error) {

	result, err := s.notesService.DeleteNote(in)
	return result, err
}
