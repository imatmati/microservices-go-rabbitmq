package main

import (
	"flag"
	"net/http"
	"payment/handler"
	"payment/logger"
)

func main() {
	output := flag.String("out", "stdout", "writer for logs")
	addr := flag.String("addr", ":8080", "server listener address")
	prefix := flag.String("prefix", "payment : ", "prefix for payment services logs")
	flag.Parse()
	logger.InitLogger(*output, *prefix)

	http.HandleFunc("/payment", handler.PaymentHandler)
	http.ListenAndServe(*addr, nil)
}
