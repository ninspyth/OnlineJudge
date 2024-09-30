package handlers

import (
	"github.com/gin-gonic/gin"
	er "github.com/ninspyth/OnlineJudge/helper"
	"github.com/streadway/amqp"
	"io"
	"log"
	"net/http"
)

// Status OK
const OK int = 200

var SenderAmqpChannel *amqp.Channel
var SenderAmqpConnection *amqp.Connection

func InitAmqp() error {

	var err error

	//Connecting to Amqp
	SenderAmqpConnection, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Error Connecting to RabbitMQ")
	}

	//Create a channel
	SenderAmqpChannel, err = SenderAmqpConnection.Channel()
	if err != nil {
		log.Fatalf("Error Creating a channel")
	}

	//Create a task queue to publish and consume
	if _, err := SenderAmqpChannel.QueueDeclare("request_queue", true, false, false, false, nil); err != nil {
		log.Fatalf("Error Creating a queue")
	}
	return nil
}

func StopAmqp() {
	SenderAmqpConnection.Close()
	SenderAmqpChannel.Close()
}

func SubmitHandle(c *gin.Context) {

	//Store the request to push to queue
	body := c.Request
	req, err := io.ReadAll(body.Body)
	er.HandleError("Failed to read body", err)

	Message := amqp.Publishing{
		Headers:         amqp.Table{},
		ContentType:     "text/plain",
		ContentEncoding: "",
		Body:            []byte(req),
		DeliveryMode:    amqp.Persistent, // 1=non-persistent, 2=persistent
	}

	//Publish the request to the Exchange
	if err := SenderAmqpChannel.Publish("", "request_queue", false, false, Message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to publish to queue"})
	} else {
		c.JSON(http.StatusAccepted, gin.H{"message": "Succesfully published message to queue"})
	}
}

func RunHandle(c *gin.Context) {
	c.String(OK, "Code Ran")
}
