package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to reset the database: %w", err)
	}
	fmt.Println("Database reset successful!")
	return nil
}
