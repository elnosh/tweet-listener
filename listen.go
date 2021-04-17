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
	return rdClient
}

func main() {
	fmt.Println("hello")
}
