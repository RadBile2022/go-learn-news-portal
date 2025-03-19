package infrastructure

import (
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/rabbitmq/amqp091-go"
)

type RabbitConfig struct {
	RabbitConnection string
}

type RabbitMq interface {
	Close()

	GetConnection() *amqp091.Connection
	GetChannel() *amqp091.Channel

	Publish(exchange string, routingKey string, body []byte) error
	PublishDelay(routingKey string, body []byte, delay int64) error
}

func NewRabbit(cfg RabbitConfig) RabbitMq {
	var err error

	rmq := &rabbitMq{
		config:       cfg,
		maxRetryConn: 3,
		connAttempts: 0,
	}

	if err = rmq.load(); err != nil {
		log.Fatalf("run: failed to init rabbitmq: %v", err)
	}

	rmq.quitChan = make(chan bool)

	go rmq.handleDisconnect()

	return rmq
}

type rabbitMq struct {
	conn         *amqp091.Connection
	channel      *amqp091.Channel
	config       RabbitConfig
	closeChan    chan *amqp091.Error
	quitChan     chan bool
	connAttempts int
	maxRetryConn int
}

func (rmq *rabbitMq) GetConnection() *amqp091.Connection {
	return rmq.conn
}

func (rmq *rabbitMq) GetChannel() *amqp091.Channel {
	return rmq.channel
}

func (rmq *rabbitMq) Close() {
	rmq.quitChan <- true
	slog.Info("shutting down rabbitMQ's connection...")
	<-rmq.quitChan
}

func (rmq *rabbitMq) Publish(exchange string, routingKey string, body []byte) error {
	return rmq.publish(exchange, routingKey, body, 0)
}

func (rmq *rabbitMq) PublishDelay(routingKey string, body []byte, delay int64) error {
	return rmq.publish("delayed", routingKey, body, delay)
}

func (rmq *rabbitMq) publish(exchange string, routingKey string, body []byte, delay int64) error {
	var err error

	headers := make(amqp091.Table)

	if delay != 0 {
		headers["x-delay"] = delay
	}

	msg := amqp091.Publishing{
		DeliveryMode: amqp091.Persistent,
		ContentType:  "application/json",
		Timestamp:    time.Now(),
		Headers:      headers,
		Body:         body,
		AppId:        "simpp-auth",
	}

	if rmq.channel.IsClosed() {
		rmq.channel, _ = rmq.conn.Channel()
	}

	err = rmq.channel.Publish(exchange, routingKey, false, false, msg)
	if err != nil {
		return err
	}

	return nil
}

func (rmq *rabbitMq) load() error {
	var err error

	rmq.conn, err = amqp091.Dial(rmq.config.RabbitConnection)
	if err != nil {
		return err
	}

	rmq.channel, err = rmq.conn.Channel()
	if err != nil {
		return err
	}

	slog.Info("connection to rabbitMQ established")

	rmq.closeChan = make(chan *amqp091.Error)
	rmq.conn.NotifyClose(rmq.closeChan)

	err = rmq.channel.ExchangeDeclare("amq.direct", "direct", true, false, false, false, nil)
	if err != nil {
		return err
	}

	args := make(amqp091.Table)
	args["x-delayed-type"] = "direct"
	err = rmq.channel.ExchangeDeclare("delayed", "x-delayed-message", true, false, false, false, args)
	if err != nil {
		return errors.Wrapf(err, "declaring exchange %q", "delayed")
	}

	return nil
}

func (rmq *rabbitMq) handleDisconnect() {
	for {
		select {
		case errChan := <-rmq.closeChan:
			if errChan != nil {
				slog.Error("rabbitMQ disconnection", slog.Any("error", errChan))
			}
		case <-rmq.quitChan:
			rmq.conn.Close()
			slog.Info("...rabbitMQ has been shut down")
			rmq.quitChan <- true
			return
		}

		if rmq.connAttempts >= rmq.maxRetryConn {
			slog.Error("max retrying connection attempts", slog.Int("attempts", rmq.connAttempts))
			os.Exit(1)
		}

		slog.Info("trying to reconnect to rabbitMQ", slog.Int("attempts", rmq.connAttempts), slog.Int("max_attempts", rmq.maxRetryConn))

		time.Sleep(5 * time.Second)

		if err := rmq.load(); err != nil {
			rmq.connAttempts += 1
			slog.Error("rabbitMQ error: %v", slog.Any("error", err))
			return
		}

		// reset connection attempts
		rmq.connAttempts = 0
	}
}
