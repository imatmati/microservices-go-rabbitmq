package messaging

import (
	"account/check/services"
	"account/logger"
	l "account/utils/language"
	"strconv"

	"github.com/streadway/amqp"
)

var PubConn *amqp.Connection
var ConsConn *amqp.Connection

const ConsQueueName = "check"

func InitMessaging(url string) {
	var err error
	ConsConn, err = amqp.Dial(url)
	l.PanicIf(err)

	PubConn, err = amqp.Dial(url)
	l.PanicIf(err)

	ch, err := ConsConn.Channel()
	l.PanicIf(err)

	err = ch.ExchangeDeclare("finance", "direct", true, false, false, false, nil)
	l.PanicIf(err)

	logger.Logger.Println("Exchange 'finance' declared in AMQP")
	_, err = ch.QueueDeclare(ConsQueueName, true, false, false, false, nil)
	l.PanicIf(err)

	logger.Logger.Println("Queue 'check' declared in AMQP")
	err = ch.QueueBind(ConsQueueName, "check", "finance", false, nil)
	l.PanicIf(err)
	logger.Logger.Println("Queue binding 'check' to 'finance'")

}

func consumeCheckRequest(ch, chResp *amqp.Channel) {
	msgs, err := ch.Consume(ConsQueueName, "", false, false, false, false, nil)
	l.PanicIf(err)
	for msg := range msgs {

		logger.Logger.Printf("Message received : %v\n", msg)
		account := string(msg.Body)
		check := services.CheckAccount(account)
		logger.Logger.Printf("Check for account %s is %s\n", account, []byte(strconv.FormatBool(check)))
		err = chResp.Publish("finance", msg.ReplyTo, false, false, amqp.Publishing{

			ContentType:   "text/plain",
			Body:          []byte(strconv.FormatBool(check)),
			CorrelationId: msg.CorrelationId,
		})
		l.PanicIf(err)
		msg.Ack(false)

	}
}

func Start() {

	chResp, err := PubConn.Channel()
	l.PanicIf(err)
	defer chResp.Close()
	ch, err := ConsConn.Channel()
	l.PanicIf(err)
	defer ch.Close()

	consumeCheckRequest(ch, chResp)
}
