package messaging

import (
	"account/check/messaging"
	"account/logger"
	l "account/utils/language"
	"account/utils/random"

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

func consumeCheckResponses(ch *amqp.Channel) <-chan amqp.Delivery {
	msgs, err := ch.Consume(ConsQueueName, "check consumer", false, false, false, false, nil)
	l.PanicIf(err)
	logger.Logger.Println("Response message consumer set")
	return msgs
}

func sendCheckRequestFor(account string, ch *amqp.Channel) string {
	corrId := random.RandomString(32)
	err := ch.Publish("finance", messaging.ConsQueueName, false, false, amqp.Publishing{
		ContentType:   "text/plain",
		Body:          []byte(account),
		ReplyTo:       ConsQueueName,
		CorrelationId: corrId,
	})
	l.PanicIf(err)
	logger.Logger.Printf("Check request message for account %s has correlation id %s and was published on %s\n", account, corrId, messaging.ConsQueueName)
	return corrId
}

func treatResponsesFromCheck(msgs <-chan amqp.Delivery, corrId string) bool {
	response := false
	for msg := range msgs {
		logger.Logger.Printf("Message received : %v\n", msg)
		if msg.CorrelationId != corrId {
			msg.Nack(false, true)
			continue
		}
		check := msg.Body
		if string(check) == "true" {
			response = true
		}
		msg.Ack(false)
		break
	}
	return response
}

func CheckAccount(account string) bool {

	// Init channels and defer close
	chResp, err := ConsConn.Channel()
	l.PanicIf(err)
	defer chResp.Close()
	ch, err := PubConn.Channel()
	l.PanicIf(err)
	defer ch.Close()

	msgs := consumeCheckResponses(chResp)
	corrId := sendCheckRequestFor(account, ch)
	return treatResponsesFromCheck(msgs, corrId)

}
