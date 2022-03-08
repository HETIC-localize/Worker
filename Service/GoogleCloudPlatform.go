package Service

// @see https://cloud.google.com/pubsub/docs/create-topic-client-libraries

import (
	"HETIC-localize/Worker/Model"
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

func GCPPubSubGetTask() Model.Task {

	task := Model.Task{}

	projectID := os.Getenv("GCP_PUBSUB_PROJECT_ID")
	subID := os.Getenv("GCP_PUBSUB_SUB_ID")

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)

	if err != nil {
		_ = fmt.Errorf("pubsub.NewClient: %v", err)
	}

	defer client.Close()

	var mu sync.Mutex
	sub := client.Subscription(subID)
	cctx, cancel := context.WithCancel(ctx)
	err = sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {

		err = json.Unmarshal([]byte(msg.Data), &task)

		if err != nil {
			println(err)
		}

		mu.Lock()
		defer mu.Unlock()
		msg.Ack()
		cancel()
	})

	if err != nil {
		_ = fmt.Errorf("pubsub.NewClient: %v", err)
	}

	return task
}
