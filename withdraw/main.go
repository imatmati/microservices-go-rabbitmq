package main

import (
	"account/logger"
	l "account/utils/language"
	"account/withdraw/handler"
	"account/withdraw/messaging"
	"flag"
	"net/http"
)

func main() {
	output := flag.String("out", "stdout", "writer for logs")
	addr := flag.String("addr", ":8080", "server listener address")
	prefix := flag.String("prefix", "withdraw : ", "prefix for withdraw services logs")
	amqp := flag.String("amqp", "amqp://guest:guest@localhost:5672/", "amqp URL")
	flag.Parse()
	logger.InitLogger(*output, *prefix)
	messaging.InitMessaging(*amqp)
	http.HandleFunc("/withdraw", handler.WithdrawHandler)
	err := http.ListenAndServe(*addr, nil)
	l.PanicIf(err)
}
