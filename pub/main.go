package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

const (
	PROJECT_ID = "pub-sub-test-392520"
	TOPIC_ID   = "fruits"
	MSG        = " I LOVE BOTH ANDY & ZAYDEN"
	KEY        = "./key.json"
)

func main() {
	publish(PROJECT_ID, TOPIC_ID, MSG, KEY)
}

func publish(projectID, topicID, msg string, key string) {

	// generate options with credentials
	opts := []option.ClientOption{}
	opts = append(opts, option.WithCredentialsFile(key))

	// get context
	ctx := context.Background()

	// create pubsub client
	pubSubClient, err := pubsub.NewClient(ctx, projectID, opts...)
	handleErr(err)

	// close client after completion
	defer pubSubClient.Close()

	// convert msg into pub sub format
	message := &pubsub.Message{
		Data: []byte(msg),
	}

	// fetch topic
	topic := pubSubClient.Topic(topicID)

	// publish on topic
	result := topic.Publish(ctx, message)

	// server generated id is returned for published msg
	serverId, err := result.Get(ctx)
	handleErr(err)
	fmt.Printf("Published a message; msg ID: %v\n", serverId)

}

func handleErr(err error) {
	if err != nil {
		log.Panic(err)
		return
	}
}
