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
{"feedUrl": "https://example.com", "feedMeta": {}, "article": "<html/>", "articleMeta": {}}
```

## Tools

### fetch
in: `<Feed>`
out: `<FetchedFeed>`

### parse
in: `<FetchedFeed>`
out: `[<FeedItem>]`

### store
in: `[<FeedItem>]`


## Pipelines

### fetch -> parse -> "store"
```command
$ echo '{"url":"http://rss.slashdot.org/Slashdot/slashdotMain"}' | go run cmd/fetch/main.go | go run cmd/parse/main.go
```