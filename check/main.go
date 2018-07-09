package main

import (
	"account/check/handler"
	"account/check/messaging"
	"account/logger"
	l "account/utils/language"
	"flag"
	"net/http"
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

	http.HandleFunc("/check", handler.CheckHandler)
	err := http.ListenAndServe(*addr, nil)
	l.PanicIf(err)
}
