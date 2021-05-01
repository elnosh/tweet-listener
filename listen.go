package main

import (
	"fmt"

	"github.com/miguelhun/tweet-listener/redis"
)

func main() {
	listenTweets()
}

func listenTweets() {
	pubsub := redis.RedisClient.Subscribe("tweetChannel")

	ch := pubsub.Channel()

	for msg := range ch {
		fmt.Printf("tweet is: %s \n", msg.Payload)
	}
}
