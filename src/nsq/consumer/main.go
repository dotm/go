package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/bitly/go-nsq"
)

type myMessageHandler struct{}

// HandleMessage implements the Handler interface.
func (h *myMessageHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
		return nil
	}

	err := processMessage(m.Body)

	// Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
	return err
}

func main() {
	// Instantiate a consumer that will subscribe to the provided channel.
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("topic", "channel", config)
	if err != nil {
		log.Fatal(err)
	}

	// Set the Handler for messages received by this Consumer. Can be called multiple times.
	// See also AddConcurrentHandlers.
	consumer.AddHandler(&myMessageHandler{})

	// Use nsqlookupd to discover nsqd instances.
	// See also ConnectToNSQD, ConnectToNSQDs, ConnectToNSQLookupds.
	err = consumer.ConnectToNSQLookupd("localhost:4161")
	if err != nil {
		log.Fatal(err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	time.AfterFunc(time.Second, func() {
		wg.Done()
	})
	wg.Wait()

	// Gracefully stop the consumer.
	consumer.Stop()
}

func processMessage(b []byte) error {
	originalMessage := string(b) //you can do JSON unmarshalling here
	fmt.Println(originalMessage)
	return nil
}
