package handler

import (
	"context"

	"github.com/sonnq2010/blog-aggregator/internal/command"
	"github.com/sonnq2010/blog-aggregator/internal/state"
)

func ResetHandler(s *state.State, c command.Command) error {
	return s.DB.DeleteAllUsers(context.Background())
}
