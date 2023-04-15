package producer

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"prac8/errorHandler"
	"time"
)

type defaultServiceProducer struct {
	connection *amqp091.Connection
	channel    *amqp091.Channel
	queueName  string
}

func NewDefault(connectionString string, queueName string) (producer *defaultServiceProducer) {
	producer = &defaultServiceProducer{}
	var err error
	producer.connection, err = amqp091.Dial(connectionString)
	errorHandler.HandleError(err, "Failed to connect to RabbitMQ")

	producer.channel, err = producer.connection.Channel()
	errorHandler.HandleError(err, "Failed to create channel")

	producer.queueName = queueName

	return producer
}

func (producer *defaultServiceProducer) SendMessage(content []byte) {
	queue, err := producer.channel.QueueDeclare(
		producer.queueName, // name
		false,              // durable
		false,              // delete when unused
		false,              // exclusive
		false,              // no-wait
		nil,                // arguments
	)
	errorHandler.HandleError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello World!"
	err = producer.channel.PublishWithContext(ctx,
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	errorHandler.HandleError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
