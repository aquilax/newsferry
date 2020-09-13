package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/aquilax/nesferry/common"
)

func main() {
	var data []byte
	var err error
	data, err = ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	var feed common.Feed
	json.Unmarshal(data, &feed)

	var resp *http.Response
	resp, err = http.Get(feed.Url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fetchedFeed := common.FetchedFeed{Feed: feed, FeedBody: string(data)}
	fetchedFeed.Feed = feed
	data, err = json.Marshal(fetchedFeed)
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(data)
}
