package services

import (
	"github.com/rabbitmq/amqp091-go"
	"log"
	"prac8/errorHandling"
)

type defaultServiceConsumer struct {
	service    Service
	connection *amqp091.Connection
	channel    *amqp091.Channel
	queueName  string
}

func NewDefault(connectionString string, queueName string) (consumer *defaultServiceConsumer) {
	consumer = &defaultServiceConsumer{}
	var err error
	consumer.connection, err = amqp091.Dial(connectionString)
	errorHandling.HandleError(err, "Failed to connect to RabbitMQ")

	consumer.channel, err = consumer.connection.Channel()
	errorHandling.HandleError(err, "Failed to create channel")

	consumer.queueName = queueName
	return consumer
}

func (consumer *defaultServiceConsumer) SetService(service Service) {
	consumer.service = service
}

func (consumer *defaultServiceConsumer) StartConsuming(handler func([]byte)) {
	queue, err := consumer.channel.QueueDeclare(
		consumer.queueName, // name
		false,              // durable
		false,              // delete when unused
		false,              // exclusive
		false,              // no-wait
		nil,                // arguments
	)
	errorHandling.HandleError(err, "Failed to declare a queue")

	msgs, err := consumer.channel.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	errorHandling.HandleError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			handler(d.Body)
			consumer.service.OnConsume(d.Body)
			d.Ack(true)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
