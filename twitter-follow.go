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

	accountsToFollow := [...]int64{
		4128569116,
	}

	for j := 0; j < len(accountsToFollow); j++ {
		// Follows users
		friendshipCreateParams := &twitter.FriendshipCreateParams{UserID: accountsToFollow[j]}
		user, _, err := client.Friendships.Create(friendshipCreateParams)
		if err != nil {
			fmt.Printf("Error following %v \n", accountsToFollow[j])
		} else {
			fmt.Printf("Followed %v \n", user.ScreenName)
		}
		if j%5 == 0 {
			time.Sleep(5 * time.Second)
		} else {
			time.Sleep(150 * time.Millisecond)
		}
	}
}
