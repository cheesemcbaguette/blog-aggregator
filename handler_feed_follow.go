package main

import (
	"context"
	"example.com/username/blog-aggregator/cheese/internal/database"
	"fmt"
	"github.com/google/uuid"
	"net/url"
	"time"
)

func handlerFollowFeed(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("please provide a feed URL")
	}

	u, err := url.ParseRequestURI(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("please provide a valid URL")
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

	// Look up feeds by URL
	feed, err := s.db.GetFeedByURL(context.Background(), u.String())

	if err != nil {
		return fmt.Errorf("couldn't find feed in the db: %w", err)
	}

	// Add the feed follow to the database
	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        feedID,
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    currentUser.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return fmt.Errorf("failed to create feed follow: %w", err)
	}

	// Print out the new feed details
	fmt.Printf("Feed follow added:\nID: %s\nFeed name: %s\nUser name: %s\n", feed.ID, feed.Name, currentUser.Name)

	return nil
}

func handlerFollowing(s *state, cmd command) error {

	// Look up feeds by URL
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), s.cfg.CurrentUserName)

	if err != nil {
		return fmt.Errorf("error searching follows for the current user: %w", err)
	}

	if len(feedFollows) == 0 {
		return fmt.Errorf("couldn't find any followed feeds for the current user: ")
	}

	for _, follows := range feedFollows {
		fmt.Printf("Feed name: %s\nFollowed by: %s\n", follows.FeedName, follows.UserName)
	}

	return nil
}
