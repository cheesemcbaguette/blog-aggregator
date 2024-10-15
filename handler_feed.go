package main

import (
	"context"
	"example.com/username/blog-aggregator/cheese/internal/database"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("please provide both feed name and URL")
	}

	// Get the current user from the config
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't find user in the db: %w", err)
	}

	// Generate a new UUID for the feed
	feedID := uuid.New()

	// Get current timestamp
	now := time.Now()

	// Add the feed to the database
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        feedID,
		CreatedAt: now,
		UpdatedAt: now,
		Name:      cmd.Args[0], // Feed name
		Url:       cmd.Args[1], // Feed URL
		UserID:    currentUser.ID,
	})

	if err != nil {
		return fmt.Errorf("failed to create feed: %w", err)
	}

	// Add the feed follow to the database
	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		FeedID:    feed.ID,
		UserID:    currentUser.ID,
	})

	if err != nil {
		return fmt.Errorf("failed to create feed follow: %w", err)
	}

	// Print out the new feed details
	fmt.Printf("Feed added:\nID: %s\nName: %s\nURL: %s\n", feed.ID, feed.Name, feed.Url)
	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get feeds from the database: %w", err)
	}

	// Check if there are no feeds
	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	// Print the feeds in the required format
	fmt.Println("Feeds:")
	for _, feed := range feeds {
		fmt.Printf("Feed: %s\nURL: %s\nAdded by: %s\n", feed.FeedName, feed.Url, feed.UserName)
	}

	return nil
}
