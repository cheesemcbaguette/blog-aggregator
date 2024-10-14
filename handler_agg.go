package main

import "context"

func handlerAggregate(s *state, cmd command) error {
	// Assuming feed URL is hardcoded for now
	feedURL := "https://www.wagslane.dev/index.xml"
	return aggregateFeed(context.Background(), s, feedURL)
}
