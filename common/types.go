package common

type Feed struct {
	Url string `json: url`
}

type FetchedFeed struct {
	FeedBody string `json: feedBody`
	Feed
}
