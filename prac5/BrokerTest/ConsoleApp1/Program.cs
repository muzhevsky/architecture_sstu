﻿using RabbitMQ.Client;
using RabbitMQ.Client.Events;
using System.Text;

var factory = new ConnectionFactory() { Uri = new Uri("amqps://egltnesj:4xbePXL1eFQHMtd-vC0NDtMnYwcJUOy9@hummingbird.rmq.cloudamqp.com/egltnesj") };

using (var connection = factory.CreateConnection())
using (var channel = connection.CreateModel())
{
    channel.QueueDeclare(queue: "MyQueue",
                     durable: false,
                     exclusive: false,
                     autoDelete: false,
                     arguments: null);

    var consumer = new EventingBasicConsumer(channel);
    consumer.Received += (model, ea) =>
    {
        var body = ea.Body.ToArray();
        var message = Encoding.UTF8.GetString(body);
        Console.WriteLine(" [x] Received {0}", message);
    };
    channel.BasicConsume(queue: "MyQueue",
                     autoAck: true,
                 consumer: consumer);

    Console.WriteLine(" Press [enter] to exit.");
    Console.ReadLine();
}