using RabbitMQ.Client;
using RabbitMQ.Client.Events;
using System;
using System.Text;
using MailKit.Net.Smtp;
using MimeKit;

var factory = new ConnectionFactory() { Uri = new Uri("amqps://qurjwfqt:1MK_DuygWJ-Wh01BvfPO6alW6Skrhhkt@hawk.rmq.cloudamqp.com/qurjwfqt") };

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
        var messageBody = Encoding.UTF8.GetString(body);
        Console.WriteLine(" [x] Received {0}", messageBody);

        // Создание и отправка сообщения на почту
        var message = new MimeMessage();
        message.From.Add(new MailboxAddress("Sender Name", ""));
        message.To.Add(new MailboxAddress("Recipient Name", ""));
        message.Subject = "Test message";
        message.Body = new TextPart("plain") { Text = "This is a test message." };

        using (var client = new SmtpClient())
        {
            client.Connect("smtp.gmail.com", 587, false);
            client.Authenticate("", "");
            client.Send(message);
            client.Disconnect(true);
        }
    };

    channel.BasicConsume(queue: "MyQueue",
                         autoAck: true,
                         consumer: consumer);

    Console.WriteLine(" Press [enter] to exit.");
    Console.ReadLine();
}