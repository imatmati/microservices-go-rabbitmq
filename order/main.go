package main

import (
	"flag"
	"net/http"
	"payment/logger"
	"payment/order/handler"
	"payment/order/messaging"
	l "payment/utils/language"
)

func main() {
	output := flag.String("out", "stdout", "writer for logs")
	addr := flag.String("addr", ":8080", "server listener address")
	prefix := flag.String("prefix", "payment : ", "prefix for payment services logs")
	amqp := flag.String("amqp", "amqp://guest:guest@localhost:5672/", "amqp URL")
	flag.Parse()
	logger.InitLogger(*output, *prefix)
	messaging.InitMessaging(*amqp)
	http.HandleFunc("/payment", handler.PaymentHandler)
	err := http.ListenAndServe(*addr, nil)
	l.PanicIf(err)
}
