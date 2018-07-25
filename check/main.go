package main

import (
	"account/check/handler"
	"account/check/messaging"
	"account/logger"
	"account/utils/env"
	l "account/utils/language"
	"net/http"
)

func main() {

	checkQueue := env.GetEnv("CHECK_QUEUE_NAME", "check")
	output := env.GetEnv("LOG_OUTPUT", "stdout")
	addr := env.GetEnv("LISTEN_ADDRESS", ":8080")
	prefix := env.GetEnv("LOG_PREFIX_MESSAGE", "check :")
	amqpURL := env.GetEnv("AMQP_URL", "amqp://guest:guest@localhost:5672/")

	logger.InitLogger(output, prefix)
	messaging.InitMessaging(amqpURL, checkQueue)
	go messaging.Start()
	http.HandleFunc("/", handler.HealthCheckHandler)
	http.HandleFunc("/check", handler.CheckHandler)
	err := http.ListenAndServe(addr, nil)
	l.PanicIf(err)
}
