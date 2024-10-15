package main

import (
	"context"
	"example.com/username/blog-aggregator/cheese/internal/database"
	"fmt"
	"github.com/google/uuid"
	"log"
	"strconv"
	"time"
)

func insertPost(db *database.Queries, feed database.Feed, item RSSItem) error {
	// Generate a new UUID for the feed
	postID := uuid.New()

	// Get current timestamp
	now := time.Now()

	// Get published date
	publishedDate, err := time.Parse(time.RFC1123Z, item.PubDate) // Use RFC1123Z or other formats as needed
	if err != nil {
		log.Printf("Couldn't parse published date %v: %v", item.PubDate, err)
		publishedDate = now // Default to current time if parsing fails
	}

	// Add post in db
	_, err = db.CreatePost(context.Background(), database.CreatePostParams{
		ID:          postID,
		CreatedAt:   now,
		UpdatedAt:   now,
		Title:       item.Title,
		Url:         item.Link,
		Description: item.Description,
		PublishedAt: publishedDate,
		FeedID:      feed.ID,
	})
	if err != nil {
		return err
	}
	return nil
}

func handlerBrowsePosts(s *state, cmd command, user database.User) error {
	// Post limit
	limit := 2
	if len(cmd.Args) == 1 {
		i, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			log.Printf("Couldn't parse post limit %v: %v", cmd.Args[0], err)
		} else {
			limit = i
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("couldn't get posts: %v", err)
	}

	if len(posts) == 0 {
		fmt.Println("User doesn't follow any feed, no posts to display")
	}

	for _, post := range posts {
		fmt.Printf("Title: %s\nURL: %s\nPublished at: %v\n\n", post.Title, post.Url, post.PublishedAt)
	}

	return nil

}
