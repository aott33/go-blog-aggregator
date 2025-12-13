package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aott33/gator/internal/database"
	"github.com/google/uuid"
)

func handlerCreateFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		return errors.New("name is required")
	}

	if len(cmd.args) == 1 {
		return errors.New("url is required")
	}
	
	ctx := context.Background()

	name := cmd.args[0]
	url := cmd.args[1]
	feed, err := s.db.CreateFeed(ctx, database.CreateFeedParams {
		ID:			uuid.New(),
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
		Name:		name,
		Url: 		url,
		UserID: 	user.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Feed: %v created in database\n", feed.Name)
	
	err = createFeedFollowForUser(ctx, s, user.ID, feed.ID)
	if err != nil {
		return err
	}	

	return nil
}

func handlerGetFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
	fmt.Println("Unable to get feeds")
		return err
	}

	for i := range feeds {
		feedName := feeds[i].Name
		feedUrl := feeds[i].Url
		feedUser := feeds[i].Name_2
		fmt.Printf("* %s\n  URL: %s\n  User: %s\n",feedName, feedUrl, feedUser)
	}

	return nil
}

func handlerFollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		return errors.New("url is required")
	}

	url := cmd.args[0]

	ctx := context.Background()

	userID := user.ID

	feed, err := s.db.GetFeed(ctx, url)
	if err != nil {
		return err
	}
	feedID := feed.ID

	err = createFeedFollowForUser(ctx, s, userID, feedID)
	if err != nil {
		return err
	}

	return nil
}

func handlerUnfollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		return errors.New("url is required")
	}

	url := cmd.args[0]

	ctx := context.Background()

	userID := user.ID

	feed, err := s.db.GetFeed(ctx, url)
	if err != nil {
		return err
	}
	feedID := feed.ID

	err = s.db.DeleteFeedFollow(ctx, database.DeleteFeedFollowParams{
		UserID: userID,
		FeedID: feedID,
	})	
	if err != nil {
		return err
	}

	return nil	
}

func handlerGetFeedFollows(s *state, cmd command, user database.User) error {
	ctx := context.Background()

	userID := user.ID

	feedFollows, err := s.db.GetFeedFollowsForUser(ctx, userID)
	if err != nil {
		return err
	}
	
	for i := range feedFollows {
		fmt.Println(feedFollows[i].FeedName)
	}

	return nil
}

func createFeedFollowForUser(ctx context.Context, s *state, userID, feedID uuid.UUID) error {
	followFeed, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID: 		uuid.New(),
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
		UserID:		userID,
		FeedID: 	feedID,
	})
	if err != nil {
		return err
	}
	
	fmt.Printf("Feed Follow Created:\n- User: %s\n- Feed: %s\n", followFeed.UserName, followFeed.FeedName)

	return nil
}

func middlewareLoggedIn(handler func(*state, command, database.User) error) func(*state, command) error {
    // return a new function
    return func(s *state, cmd command) error {
        user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
    }
}
