# newsferry

An RSS thing

## Entities

### Feed:
```json
{"feedUrl": "https://example.com", "feedMeta": {}}
```

### FetchedFeed
```json
{"feedUrl": "https://example.com", "feedMeta": {}, "feedBody": "<xml/>"}
```

### FeedItem
```json
{"feedUrl": "https://example.com", "feedMeta": {}, "article": "<html>", "articleMeta": {}}
```

## Tools

### fetcher
in: <Feed>
out: <FetchedFeed>

### parser
in: <FetchedFeed>
out: [<FeedItem>]

### storage-in
in: [<FeedItem>]


### Fetcher pipeline
```command
$ echo '{"url":"http://rss.slashdot.org/Slashdot/slashdotMain"}' | go run cmd/fetcher/fetcher.go | go run cmd/parser/parser.go
```