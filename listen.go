package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/miguelhun/tweet-listener/redis"
)

func main() {
	http.HandleFunc("/tweetStream", printTweets)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func printTweets(w http.ResponseWriter, r *http.Request) {

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	pubsub := redis.RedisClient.Subscribe("tweetChannel")

	ch := pubsub.Channel()

	for msg := range ch {
		if err = conn.WriteMessage(websocket.TextMessage, []byte(msg.Payload)); err != nil {
			log.Println(err)
		}
	}

}
