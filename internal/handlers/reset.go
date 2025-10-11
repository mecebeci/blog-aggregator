package handlers

import (
	"context"
	"fmt"
	"os"

	"github.com/mecebeci/blog-aggregator/internal/command"
	"github.com/mecebeci/blog-aggregator/internal/state"
)

func HandleReset(s *state.State, cmd command.Command) error {
	err := s.DB.DeleteAllUsers(context.Background())
	if err != nil {
		fmt.Fprintln(os.Stderr, "Reset failed: ", err)
		return err
	}

	fmt.Println("All users deleted successfully.")
	return nil
}
