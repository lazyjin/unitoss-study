package common

import (
	"github.com/streadway/amqp"
	"strconv"
)

type Rabbit struct {
	Host string
	Port int
	user string
	pw   string

	conn  *amqp.Connection
	ch    *amqp.Channel
	queue amqp.Queue
}

type RabbitMgr interface {
	ConnectRabbit(host string, port int, id string, pw string)
}

var _ RabbitMgr = &Rabbit{}

func NewRabbitManager() RabbitMgr {
	rmgr := &Rabbit{}

	return rmgr
}

func (r *Rabbit) ConnectRabbit(host string, port int, id string, pw string) {
	log.Infof("host: %v || port: %v", host, port)
	r.Host = host
	r.Port = port
	r.user = id
	r.pw = pw

	var err error

	r.conn, err = amqp.Dial("amqp://" + id + ":" + pw + "@" + r.Host + ":" + strconv.Itoa(r.Port) + "/")
	CheckErrPanic(err)

	r.ch, err = r.conn.Channel()
	CheckErrPanic(err)

	log.Info("Successfully connect to RabbitMQ...")
}

func (r *Rabbit) CloseConnRabbit() {
	r.conn.Close()
	log.Info("Successfully close RabbitMQ connection...")
}

func (r *Rabbit) CloseChanRabbit() {
	r.ch.Close()
	log.Info("Successfully close RabbitMQ channel...")
}

func (r *Rabbit) TaskQueueDeclare() {
	var err error

	r.queue, err = r.ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	CheckErrPanic(err)
}

func PublishToQueue() {

}
