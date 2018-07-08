package messaging

import (
	"payment/check/messaging"
	"payment/logger"
	l "payment/utils/language"
	"payment/utils/random"

	"github.com/streadway/amqp"
)

var PubConn *amqp.Connection
var ConsConn *amqp.Connection

const ConsQueueName = "order_check"

func InitMessaging(url string) {
	var err error
	PubConn, err = amqp.Dial(url)
	l.PanicIf(err)

	ConsConn, err = amqp.Dial(url)
	l.PanicIf(err)

	ch, err := PubConn.Channel()
	l.PanicIf(err)

	err = ch.ExchangeDeclare("finance", "direct", true, false, false, false, nil)
	l.PanicIf(err)
	logger.Logger.Println("Exchange 'finance' declared in AMQP")

	_, err = ch.QueueDeclare(ConsQueueName, true, false, false, false, nil)
	l.PanicIf(err)

	err = ch.QueueBind(ConsQueueName, ConsQueueName, "finance", false, nil)
	l.PanicIf(err)

}

func CheckAccount(account string) bool {

	chResp, err := PubConn.Channel()
	l.PanicIf(err)
	defer chResp.Close()
	ch, err := ConsConn.Channel()
	l.PanicIf(err)
	defer ch.Close()
	msgs, err := chResp.Consume(ConsQueueName, "check consumer", false, false, false, false, nil)
	l.PanicIf(err)

	corrId := random.RandomString(32)
	logger.Logger.Printf("Message correlation id : %s published on %s\n", corrId, messaging.ConsQueueName)
	err = ch.Publish("finance", messaging.ConsQueueName, false, false, amqp.Publishing{
		ContentType:   "text/plain",
		Body:          []byte(account),
		ReplyTo:       ConsQueueName,
		CorrelationId: corrId,
	})
	l.PanicIf(err)
	response := false
	for msg := range msgs {
		logger.Logger.Printf("Message received : %v\n", msg)
		if msg.CorrelationId != corrId {
			msg.Nack(false, true)
			continue
		}
		check := msg.Body
		logger.Logger.Printf("Response for %s (%s) , %s : %v", account, msg.CorrelationId, check, msg)
		msg.Ack(false)
		if string(check) == "true" {
			response = true
		}
		break
	}
	//TODO Manage the case where the loop is not executed.
	return response
}
