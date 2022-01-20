package unfolllow

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
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(false),
	}

	client := getClient(&creds)

	// Gets authenticated user
	me, _, _ := client.Accounts.VerifyCredentials(verifyParams)

	// Gets followed users' ID
	friendIDParams := &twitter.FriendIDParams{UserID: me.ID}
	friendsIds, _, err := client.Friends.IDs(friendIDParams)
	if err != nil {
		fmt.Println("Fetching user's followings failed!")
	} else {
		fmt.Printf("%v \n", len(friendsIds.IDs))
		for i := 0; i < len(friendsIds.IDs); i++ {
			// Unfollows users
			friendshipDestroyParams := &twitter.FriendshipDestroyParams{UserID: friendsIds.IDs[i]}
			user, _, err := client.Friendships.Destroy(friendshipDestroyParams)
			if err != nil {
				fmt.Printf("Error unfollowing %v \n", friendsIds.IDs[i])
			} else {
				fmt.Printf("Unfollowed %v \n", user.ScreenName)
			}
			if i%5 == 0 {
				time.Sleep(5 * time.Second)
			} else {
				time.Sleep(200 * time.Millisecond)
			}
		}
	}
}
