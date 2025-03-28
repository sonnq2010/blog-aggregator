package handler

import (
	"context"
	"fmt"

	"github.com/sonnq2010/blog-aggregator/internal/command"
	"github.com/sonnq2010/blog-aggregator/internal/state"
)

func UsersHandler(s *state.State, c command.Command) error {
	users, err := s.DB.GetAllUsers(context.Background())
	if err != nil {
		return err
	}

	currentUser := s.Config.CurrentUserName
	for _, user := range users {
		str := fmt.Sprintf("* %s ", user.Name)
		if currentUser == user.Name {
			str += "(current)"
		}
		fmt.Println(str)
	}

	return nil
}
