package common

import (
	"time"

	"github.com/mmcdole/gofeed"
)

type FeedRequest struct {
	Url string `json:"url"`
}

type FetchedFeed struct {
	FeedBody    string      `json:"feedBody"`
	FeedRequest FeedRequest `json:"feedRequest"`
}

type FeedConfig struct {
	Title           string            `json:"title,omitempty"`
	Description     string            `json:"description,omitempty"`
	Link            string            `json:"link,omitempty"`
	FeedLink        string            `json:"feedLink,omitempty"`
	Updated         string            `json:"updated,omitempty"`
	UpdatedParsed   *time.Time        `json:"updatedParsed,omitempty"`
	Published       string            `json:"published,omitempty"`
	PublishedParsed *time.Time        `json:"publishedParsed,omitempty"`
	Author          *gofeed.Person    `json:"author,omitempty"`
	Language        string            `json:"language,omitempty"`
	Image           *gofeed.Image     `json:"image,omitempty"`
	Copyright       string            `json:"copyright,omitempty"`
	Generator       string            `json:"generator,omitempty"`
	Categories      []string          `json:"categories,omitempty"`
	Custom          map[string]string `json:"custom,omitempty"`
	FeedType        string            `json:"feedType"`
	FeedVersion     string            `json:"feedVersion"`
}

type FeedItem struct {
	FeedRequest FeedRequest `json:"feedRequest"`
	FeedConfig  FeedConfig  `json:"feedConfig"`
	Item        gofeed.Item `json:"feedItem"`
}
