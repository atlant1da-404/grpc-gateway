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
