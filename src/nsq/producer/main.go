package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/bitly/go-nsq"
)

func main() {
	producer := createProducer()

	messageBody := []byte(getUserInput())
	topicName := "topic"
	publish(producer, topicName, messageBody)

	// Gracefully stop the producer.
	producer.Stop()
}

func createProducer() *nsq.Producer {
	// Instantiate a producer.
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	}

	return producer
}

func publish(producer *nsq.Producer, topicName string, messageBody []byte) {
	// Synchronously publish a single message to the specified topic.
	// Messages can also be sent asynchronously and/or in batches.
	err := producer.Publish(topicName, messageBody)
	if err != nil {
		log.Fatal(err)
	}
}

func getUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	for scanner.Scan() {
		input = scanner.Text()
		break //read one line only
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return input
}
