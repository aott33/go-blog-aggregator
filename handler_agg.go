package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/aott33/gator/internal/database"
	"github.com/google/uuid"
)

 func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("duration command needed - eg. 1s, 1m, 1h")
	}
	
	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			fmt.Println("scrape error:", err)
		}
	}

	return nil
}

func handlerBrowse(s *state, cmd command, user database.User) error {
	ctx := context.Background()

	limit := 2

	if len(cmd.args) == 1 {
		i, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return err
		}
		if i <= 0 {
       		return fmt.Errorf("limit must be a positive number")
    	}
		limit = i
	}
	

	posts, err := s.db.GetPosts(ctx, database.GetPostsParams{
		UserID: user.ID,
		Limit: int32(limit),
	})
	if err != nil {
		return err
	}

	for _, p := range posts {
 		fmt.Println("Title:", p.Title)
    	fmt.Println("URL:  ", p.Url)
    	if p.Description.Valid {
        	fmt.Println("Desc:", p.Description.String)
    	}
    	if p.PublishedAt.Valid {
        	fmt.Println("Published:", p.PublishedAt.Time.Format(time.RFC1123))
    	}
    	fmt.Println("-----")
	}

	return nil
}

func scrapeFeeds(s *state) error {
	ctx := context.Background()
	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return err
	}

	feedMarked, err := s.db.MarkFeedFetched(ctx, feed.ID)
	if err != nil {
		return err
	}

	feedFetch, err := fetchFeed(ctx, feedMarked.Url)
	if err != nil {
		return err
	}
	
	fmt.Println(feedFetch.Channel.Title)

	for i := range feedFetch.Channel.Item {
		desc := toNullString(feedFetch.Channel.Item[i].Description)

		pubTime := parseFlexibleTime(feedFetch.Channel.Item[i].PubDate)

		_, err := s.db.CreatePosts(ctx, database.CreatePostsParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: feedFetch.Channel.Item[i].Title,
			Url: feedFetch.Channel.Item[i].Link,
			Description: desc,
			PublishedAt: pubTime,
			FeedID: feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
        		// ignore duplicate URL
        		continue
    		}
			fmt.Println("error creating post:", err)
		}
	}


	return nil
}

func parseFlexibleTime(input string) sql.NullTime {
	if input == "" {
		return sql.NullTime{}
	}

	formats := []string{
		time.RFC3339,
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
		time.RFC822,
		"Mon, 02 Jan 2006 15:04:05 -0700",
	}

	for _, format := range formats {
		if t, err := time.Parse(format, input); err == nil {
			return sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}
	}
	
	fmt.Printf("could not parse pubDate %q with known formats\n", input)
	return sql.NullTime{}
}

func toNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}
	return sql.NullString{String: s, Valid: true}
}
