package main

import (
	"database/sql"
	"time"

	"github.com/ArrayOfLilly/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string	`json:"api_key"`
}

func databaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID 	 `json:"id"`
	Name      string    	 `json:"name"`
	Url       string		 `json:"url"`
	UserID    uuid.UUID 	 `json:"user_id"`
	CreatedAt time.Time 	 `json:"created_at"`
	UpdatedAt time.Time 	 `json:"updated_at"`
	LastFetchedAt *time.Time `json:"last_fetched_at"`
}

func databaseFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID: feed.ID,
		Name: feed.Name,
		Url: feed.Url,
		UserID: feed.UserID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		LastFetchedAt: nullTimeToTimePtr(feed.LastFetchedAt),
	}
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID	`json:"feed_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func databaseFeedFollowToFeedFollow(feed_follow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID: feed_follow.ID,
		UserID: feed_follow.UserID,
		FeedID: feed_follow.FeedID,
		CreatedAt: feed_follow.CreatedAt,
		UpdatedAt: feed_follow.UpdatedAt,
	}
}

func nullTimeToTimePtr(t sql.NullTime) *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}