package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/aquilax/newsferry/common"
)

func main() {
	var data []byte
	var err error
	data, err = ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	var feedRequest common.FeedRequest
	json.Unmarshal(data, &feedRequest)

	var resp *http.Response
	resp, err = http.Get(feedRequest.Url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fetchedFeed := common.FetchedFeed{FeedRequest: feedRequest, FeedBody: string(data)}
	data, err = json.Marshal(fetchedFeed)
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(data)
}
