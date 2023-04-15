package consumer

import (
	"bytes"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"prac8/errorHandler"
	"time"
)

type defaultServiceConsumer struct {
	connection *amqp091.Connection
	channel    *amqp091.Channel
	queueName  string
}

func NewDefault(connectionString string, queuerName string) (consumer *defaultServiceConsumer) {
	consumer = &defaultServiceConsumer{}
	var err error
	consumer.connection, err = amqp091.Dial(connectionString)
	errorHandler.HandleError(err, "Failed to connect to RabbitMQ")

	consumer.channel, err = consumer.connection.Channel()
	errorHandler.HandleError(err, "Failed to create channel")

	consumer.queueName = queuerName
	return consumer
}

func (consumer *defaultServiceConsumer) StartConsuming() {
	queue, err := consumer.channel.QueueDeclare(
		consumer.queueName, // name
		false,              // durable
		false,              // delete when unused
		false,              // exclusive
		false,              // no-wait
		nil,                // arguments
	)
	errorHandler.HandleError(err, "Failed to declare a queue")

	msgs, err := consumer.channel.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	errorHandler.HandleError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			log.Printf("Done")
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
