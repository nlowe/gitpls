package main

import (
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	consumerKey, found := os.LookupEnv("TWITTER_KEY")
	if !found {
		panic("Missing required key: TWITTER_KEY")
	}

	consumerSecret, found := os.LookupEnv("TWITTER_SECRET")
	if !found {
		panic("Missing required key: TWITTER_SECRET")
	}

	accessToken, found := os.LookupEnv("TWITTER_ACCESS_TOKEN")
	if !found {
		panic("Missing required key: TWITTER_ACCESS_TOKEN")
	}

	accessSecret, found := os.LookupEnv("TWITTER_ACCESS_SECRET")
	if !found {
		panic("Missing required key: TWITTER_ACCESS_SECRET")
	}

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	client.Statuses.Update("It works!", nil)
}
