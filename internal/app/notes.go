package app

import (
	"context"
	"grpc-gateway/pkg/gen/notes"
)

type Notes struct {
	notes.UnimplementedNotesServer
}

func NewNotes() *Notes {
	return &Notes{}
}

func (s *Notes) GetNote(ctx context.Context, in *notes.NotesRequest) (*notes.NotesResponse, error) {

	switch in.Id {
	case 1:
		return &notes.NotesResponse{
			Name:  "atlant1da-404",
			Color: "Black",
			Count: 3,
		}, nil
	case 2:
		return &notes.NotesResponse{
			Name:  "bob",
			Color: "Red",
			Count: 1,
		}, nil
	}

	return nil, nil
}
