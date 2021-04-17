package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var (
	rdClient *redis.Client
)

func rdsClient() *redis.Client {
	rdClient = redis.NewClient(&redis.Options{})
	// ask about this client return
	return rdClient
}

func main() {
	rdsClient()
	listenTweets()
}

func listenTweets() {
	pubsub := rdClient.Subscribe("tweetChannel")

	ch := pubsub.Channel()

	for msg := range ch {
		fmt.Printf("tweet is: %s \n", msg.Payload)
	}
}
