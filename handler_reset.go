package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	ctx := context.Background()

	if err := s.db.ResetFeedFollows(ctx); err != nil {
		fmt.Println("Unsuccessful!")
		return err
	}
	if err := s.db.ResetFeeds(ctx); err != nil {
		fmt.Println("Unsuccessful!")
		return err
	}
	if err := s.db.ResetUsers(ctx); err != nil {
		fmt.Println("Unsuccessful!")
		return err
	}

	fmt.Println("Successful: users, feeds, and feed_follows reset")
	return nil
}
