package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aquilax/newsferry/common"
)

type DebugStorage struct{}

func (ds DebugStorage) Set(feedItem common.FeedItem) error {
	_, err := fmt.Printf("%+v\n", feedItem)
	return err
}

func main() {
	var err error
	scanner := bufio.NewScanner(os.Stdin)

	var ds = DebugStorage{}

	var feedItem common.FeedItem
	for scanner.Scan() {
		err = json.Unmarshal(scanner.Bytes(), &feedItem)
		if err != nil {
			panic(err)
		}
		err = ds.Set(feedItem)
		if err != nil {
			panic(err)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
