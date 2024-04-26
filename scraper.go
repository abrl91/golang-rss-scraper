package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/abrl91/golang-rss-scraper/internal/database"
	"github.com/google/uuid"
)

func startScraping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {
	log.Printf("Starting scraping on %v goroutines every %s duration", concurrency, timeBetweenRequest)

	ticker := time.NewTicker(timeBetweenRequest)

	// the ; ; <-ticker.C is a way to make it run immediately at first and then run every timeBetweenRequest
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Println("Error getting feeds to fetch", err)
			continue
		}

		wg := &sync.WaitGroup{}

		for _, feed := range feeds {
			wg.Add(1) // increments the WaitGroup counter by 1
			go scrapeFeed(db, wg, feed)
		}
		wg.Wait() // blocks until the WaitGroup counter is 0
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done() // decrements the counter by 1

	log.Printf("Scraping feed %s", feed.Url)
	_, err := db.UpdateFeedLastFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Error updating feed last fetched %s", feed.Url)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Printf("Error fetching feed %s", feed.Url)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		t, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Error parsing date %s", item.PubDate)
			continue
		}

		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Body:        description,
			PublishedAt: t,
			Url:         item.Link,
			FeedID:      feed.ID,
		})

		if err != nil {
			if strings.Contains(err.Error(), "unique constraint") {
				continue
			}
			log.Printf("Error creating post %s", item.Title)
		}
	}
	log.Printf("Feed %s scraped, %v posts found", feed.Title, len(rssFeed.Channel.Item))
}
