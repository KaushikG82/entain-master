package service

import (
	"git.neds.sh/matty/entain/sport/db"
	"git.neds.sh/matty/entain/sport/proto/sport"
	"golang.org/x/net/context"
)

type Sport interface {
	// ListEvents will return a collection of sport events.
	ListEvents(ctx context.Context, in *sport.ListEventRequest) (*sport.ListEventsResponse, error)
}

// sportService implements the Sport interface.
type sportService struct {
	sportsRepo db.SportRepo
}

// NewSportService instantiates and returns a new sportService.
func NewSportService(sportsRepo db.SportRepo) Sport {
	return &sportService{sportsRepo}
}

func (s *sportService) ListEvents(ctx context.Context, in *sport.ListEventRequest) (*sport.ListEventsResponse, error) {
	events, err := s.sportsRepo.List(in.Filter)
	if err != nil {
		return nil, err
	}

	return &sport.ListEventsResponse{Events: events}, nil
}
