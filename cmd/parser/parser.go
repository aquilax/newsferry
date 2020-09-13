package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/aquilax/newsferry/common"
	"github.com/mmcdole/gofeed"
)

func main() {
	var data []byte
	var err error
	data, err = ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	var fetchedFeed common.FetchedFeed
	json.Unmarshal(data, &fetchedFeed)

	fp := gofeed.NewParser()

	feed, err := fp.ParseString(fetchedFeed.FeedBody)
	if err != nil {
		panic(err)
	}

	var feedItem common.FeedItem
	var feedConfig = common.FeedConfig{
		Title:           feed.Title,
		Description:     feed.Description,
		Link:            feed.Link,
		FeedLink:        feed.FeedLink,
		UpdatedParsed:   feed.UpdatedParsed,
		PublishedParsed: feed.PublishedParsed,
		Author:          feed.Author,
		Language:        feed.Language,
		Image:           feed.Image,
		Copyright:       feed.Copyright,
		Generator:       feed.Generator,
		Categories:      feed.Categories,
		Custom:          feed.Custom,
		FeedType:        feed.FeedType,
		FeedVersion:     feed.FeedVersion,
	}
	for _, item := range feed.Items {
		feedItem.Item = *item
		feedItem.FeedConfig = feedConfig
		feedItem.FeedRequest = fetchedFeed.FeedRequest
		data, err = json.Marshal(feedItem)
		os.Stdout.Write(data)
		os.Stdout.WriteString("\n")
	}
}
