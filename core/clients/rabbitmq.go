package clients

import (
	"context"
	"log"
	"sync"
	"time"

	"awesome.fintech.org/core/constants"
	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	connection *amqp091.Connection
}

type PublishArgs struct {
	Queue   string
	Content string
}

type SubscribeArgs struct {
	Queue    string
	Callback func(content string)
}

func (r *RabbitMQ) Publish(args PublishArgs) error {
	ch, err := r.connection.Channel()
	if err != nil {
		log.Println("[x]", err)

		return err
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		args.Queue, // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		log.Println("[x]", err)

		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp091.Publishing{
			DeliveryMode: amqp091.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(args.Content),
		})
	if err != nil {
		log.Println("[x]", err)

		return err
	}

	log.Printf("[x] Sent %s", args.Content)

	return nil
}

func (r *RabbitMQ) Subscribe(args SubscribeArgs) error {
	ch, err := r.connection.Channel()
	if err != nil {
		log.Println("[x]", err)

		return err
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		args.Queue, // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		log.Println("[x]", err)

		return err
	}

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		log.Println("[x]", err)

		return err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Println("[x]", err)

		return err
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("[*] Received a message: %s", d.Body)

			args.Callback(string(d.Body))

			log.Println("[*] Done")

			d.Ack(false)
		}
	}()

	log.Println("[*] Waiting for messages. To exit press CTRL+C")
	<-forever

	return nil
}

func (r *RabbitMQ) SubscribeWithWorkers(workers int, args SubscribeArgs) {
	var wg sync.WaitGroup

	log.Printf("[x] Spawning %d workers", workers)

	for i := 1; i <= workers; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			r.Subscribe(args)
		}()

		log.Printf("[x] Worker %d just joined the party", i)
	}

	wg.Wait()
}

func NewRabbitMQ(connection *amqp091.Connection) (r *RabbitMQ) {
	return &RabbitMQ{connection: connection}
}

func NewRabbitMQConnection(env *constants.Env) *amqp091.Connection {
	connection, err := amqp091.Dial(env.RABBITMQ_URL)

	if err != nil {
		log.Printf("[x] %s: while connecting", err)
	}

	log.Println("[x] connection established")

	return connection
}
