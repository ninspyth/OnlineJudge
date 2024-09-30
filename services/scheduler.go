package scheduler

import (
	"fmt"
	"log"
	"sync"

	"github.com/ninspyth/OnlineJudge/helper"
	"github.com/streadway/amqp"
)

const MAX_WORKERS = 5

var (
	ReceiverAmqpChannel    *amqp.Channel
	ReceiverAmqpConnection *amqp.Connection
	err                    error
)

func StopWorker() {

}

// Create Multiple Workers
func StartMultipleWorkers() {
	var wg sync.WaitGroup
	wg.Add(MAX_WORKERS)

	//Create a connection to RabbitMQ
	ReceiverAmqpConnection, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		er.HandleError("Failed to dial", err)
	}

	//creating the workers
	for i := 0; i < MAX_WORKERS; i++ {
		go func() {
			defer wg.Done()
			Worker(i, ReceiverAmqpConnection, &wg)
		}()
	}
	wg.Wait()
	log.Println("All the workers have completed the tasks")
}

func Worker(id int, conn *amqp.Connection, wq *sync.WaitGroup) {

	//Create a channel
	ReceiverAmqpChannel, err = conn.Channel()
	er.HandleError("Failed to create a channel", err)

	//Consume the message from queue
	messages, err := ReceiverAmqpChannel.Consume("request_queue", "", true, false, false, false, nil)
	if err != nil {
		er.HandleError("Failed to consume from queue", err)
	} else {
		//read the message from the queue
		for msg := range messages {
			go processMessage(id, msg)
		}
	}
}

func processMessage(id int, msg amqp.Delivery) {
	fmt.Println("Body of the message received for id: ", id, "\n", string(msg.Body))
}
