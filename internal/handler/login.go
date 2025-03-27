package handler

import (
	"errors"
	"fmt"

	"github.com/sonnq2010/blog-aggregator/internal/command"
	"github.com/sonnq2010/blog-aggregator/internal/state"
)

func LoginHandler(s *state.State, c command.Command) error {
	if len(c.Args) == 0 {
		return errors.New("username is required")
	}

	username := c.Args[0]
	if username == "" {
		return errors.New("invalid username")
	}

	err := s.Config.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Println("Logged in as", username)
	return nil
}
