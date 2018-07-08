package main

import (
	"flag"
	"net/http"
	"payment/check/handler"
	"payment/check/messaging"
	"payment/logger"
	l "payment/utils/language"
)

func main() {
	output := flag.String("out", "stdout", "writer for logs")
	addr := flag.String("addr", ":8080", "server listener address")
	prefix := flag.String("prefix", "check : ", "prefix for check services logs")
	amqp := flag.String("amqp", "amqp://guest:guest@localhost:5672/", "amqp URL")
	flag.Parse()
	logger.InitLogger(*output, *prefix)
	messaging.InitMessaging(*amqp)
	go messaging.Start()

	http.HandleFunc("/check", handler.PaymentHandler)
	err := http.ListenAndServe(*addr, nil)
	l.PanicIf(err)
}
