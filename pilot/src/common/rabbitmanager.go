package common

import (
	"github.com/streadway/amqp"
	"strconv"
)

type QueMsg <-chan amqp.Delivery

type Rabbit struct {
	Host string
	Port int
	user string
	pw   string

	conn *amqp.Connection
	ch   *amqp.Channel
	sque amqp.Queue
	rque amqp.Queue
}

type RabbitMgr interface {
	ConnectRabbit(host string, port int, id string, pw string)
	UdrSendQueueDeclare(qn string)
	ReqRecvQueueDeclare(qn string)
	PublishToQueue(msg string) error
	ConsumeQueue() (QueMsg, error)
	CloseChanRabbit()
	CloseConnRabbit()
}

var _ RabbitMgr = &Rabbit{}

func NewRabbitManager() RabbitMgr {
	rmgr := &Rabbit{}

	return rmgr
}

func (r *Rabbit) ConnectRabbit(host string, port int, id string, pw string) {
	r.Host = host
	r.Port = port
	r.user = id
	r.pw = pw

	var err error

	r.conn, err = amqp.Dial("amqp://" + id + ":" + pw + "@" + r.Host + ":" + strconv.Itoa(r.Port) + "/")
	if err != nil {
		log.Panic(err)
	}

	r.ch, err = r.conn.Channel()
	if err != nil {
		log.Panic(err)
	}

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

func (r *Rabbit) UdrSendQueueDeclare(qn string) {
	var err error

	r.sque, err = r.ch.QueueDeclare(
		qn,    // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Panic(err)
	}
}

func (r *Rabbit) ReqRecvQueueDeclare(qn string) {
	var err error

	r.rque, err = r.ch.QueueDeclare(
		qn,    // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Panic(err)
	}
}

func (r *Rabbit) PublishToQueue(msg string) error {
	err := r.ch.Publish(
		"",          // exchange
		r.sque.Name, // routing key
		false,       // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(msg),
		})
	if err != nil {
		return err
	}

	return nil
}

func (r *Rabbit) ConsumeQueue() (QueMsg, error) {
	msgs, err := r.ch.Consume(
		r.rque.Name, // queue
		"",          // consumer
		false,       // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}

func ResponseAck(d amqp.Delivery, pCh chan bool) {
	result := <-pCh

	if result {
		d.Ack(false)
	} else {
		d.Reject(true)
	}
}
