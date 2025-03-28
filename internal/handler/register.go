package handler

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sonnq2010/blog-aggregator/internal/command"
	"github.com/sonnq2010/blog-aggregator/internal/database"
	"github.com/sonnq2010/blog-aggregator/internal/state"
)

func RegisterHandler(s *state.State, c command.Command) error {
	if len(c.Args) == 0 {
		return errors.New("username is required")
	}

	username := c.Args[0]
	if username == "" {
		return errors.New("invalid username")
	}

	user, err := s.DB.CreateUser(
		context.Background(),
		database.CreateUserParams{
			ID:        uuid.UUID{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      username,
		},
	)
	if err != nil {
		return err
	}

	err = s.Config.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Println("User registered:", user.Name)
	fmt.Println("Current user:", user)

	return nil
}
