package main

import (
	"fmt"
	"os"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func getClient(creds *Credentials) *twitter.Client {
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	return client
}

func main() {
	creds := Credentials{
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
	}

	client := getClient(&creds)

	params := &twitter.UserTimelineParams{
		ScreenName:      "twitterhandle",
		TrimUser:        twitter.Bool(true),
		ExcludeReplies:  twitter.Bool(true),
		IncludeRetweets: twitter.Bool(true),
	}

	for {

		tweets, _, err := client.Timelines.UserTimeline(params)

		if err != nil {
			fmt.Printf("Fetch user timeline failed!\n")
			break
		}

		fmt.Printf("Tweet count: %v\n", len(tweets))

		for j := 0; j < len(tweets); j++ {
			if tweets[j].Retweeted {
				params := &twitter.StatusUnretweetParams{TrimUser: twitter.Bool(true)}
				_, _, err := client.Statuses.Unretweet(tweets[j].ID, params)

				if err != nil {
					fmt.Printf("Unretweet failed!\n")
				} else {
					fmt.Printf("Unretweet successful\n")
				}
			} else {
				params := &twitter.StatusDestroyParams{TrimUser: twitter.Bool(true)}
				_, _, err := client.Statuses.Destroy(tweets[j].ID, params)

				if err != nil {
					fmt.Printf("Tweet destroy failed!\n")
				} else {
					fmt.Printf("Tweet destroy successful\n")
				}
			}
			if j%5 == 0 {
				time.Sleep(4 * time.Second)
			} else {
				time.Sleep(150 * time.Millisecond)
			}
		}
	}
}
