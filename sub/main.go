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
	KEY        = "./key.json"
	SUBSC_NAME = "fruits-subs-1"
)

func main() {
	subscribe(PROJECT_ID, TOPIC_ID, KEY)
}

func subscribe(projectID, topicID, key string) {

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

	// get subscription
	subscription := pubSubClient.Subscription(SUBSC_NAME)
	ctx, cancel := context.WithCancel(ctx)
	err = subscription.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		m.Ack()
		fmt.Println("Received: ", string(m.Data))
	})
	handleErr(err)
	fmt.Println(cancel)
}

func handleErr(err error) {
	if err != nil {
		log.Panic(err)
		return
	}
}
