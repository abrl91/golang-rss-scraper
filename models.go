package main

import "github.com/abrl91/golang-rss-scraper/internal/database"

type User struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Name      string `json:"name"`
	ApiKey    string `json:"api_key"`
}

type Feed struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Title     string `json:"title"`
	Url       string `json:"url"`
	UserID    string `json:"user_id"`
}

type FeedFollow struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	UserID    string `json:"user_id"`
	FeedID    string `json:"feed_id"`
}

type Post struct {
	ID          string  `json:"id"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	Title       string  `json:"title"`
	Body        *string `json:"body"` // because it can be null, we use a pointer so that if it returns null, it will return null in the json response
	PublishedAt string  `json:"published_at"`
	Url         string  `json:"url"`
	FeedID      string  `json:"feed_id"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID.String(),
		CreatedAt: dbUser.CreatedAt.String(),
		UpdatedAt: dbUser.UpdatedAt.String(),
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID.String(),
		CreatedAt: dbFeed.CreatedAt.String(),
		UpdatedAt: dbFeed.UpdatedAt.String(),
		Title:     dbFeed.Title,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID.String(),
	}
}

func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := make([]Feed, 0, len(dbFeeds))
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}
	return feeds
}

func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.ID.String(),
		CreatedAt: dbFeedFollow.CreatedAt.String(),
		UpdatedAt: dbFeedFollow.UpdatedAt.String(),
		UserID:    dbFeedFollow.UserID.String(),
		FeedID:    dbFeedFollow.FeedID.String(),
	}
}

func databaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
	feedFollows := make([]FeedFollow, 0, len(dbFeedFollows))
	for _, dbFeedFollow := range dbFeedFollows {
		feedFollows = append(feedFollows, databaseFeedFollowToFeedFollow(dbFeedFollow))
	}
	return feedFollows
}

func databasePostToPost(dbPost database.Post) Post {
	var description *string
	if dbPost.Body.Valid {
		description = &dbPost.Body.String
	}
	return Post{
		ID:          dbPost.ID.String(),
		CreatedAt:   dbPost.CreatedAt.String(),
		UpdatedAt:   dbPost.UpdatedAt.String(),
		Title:       dbPost.Title,
		Body:        description,
		PublishedAt: dbPost.PublishedAt.String(),
		Url:         dbPost.Url,
		FeedID:      dbPost.FeedID.String(),
	}
}

func databasePostsToPosts(dbPosts []database.Post) []Post {
	posts := make([]Post, 0, len(dbPosts))
	for _, dbPost := range dbPosts {
		posts = append(posts, databasePostToPost(dbPost))
	}
	return posts
}
