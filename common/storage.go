package common

type Storage interface {
	Set(FeedItem) error
}
