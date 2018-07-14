package main

import (
	"account/logger"
	"account/utils/env"
	l "account/utils/language"
	"account/withdraw/handler"
	"account/withdraw/messaging"
	"net/http"
)

func main() {
	checkQueue := env.GetEnv("CHECK_QUEUE_NAME", "check")
	withdrawQueue := env.GetEnv("WITHDRAW_QUEUE_NAME", "check")
	output := env.GetEnv("LOG_OUTPUT", "stdout")
	addr := env.GetEnv("LISTEN_ADDRESS", ":8080")
	prefix := env.GetEnv("LOG_PREFIX_MESSAGE", "withdraw :")
	amqpURL := env.GetEnv("AMQP_URL", "amqp://guest:guest@localhost:5672/")

	logger.InitLogger(output, prefix)
	messaging.InitMessaging(amqpURL, checkQueue, withdrawQueue)
	http.HandleFunc("/withdraw", handler.WithdrawHandler)
	err := http.ListenAndServe(addr, nil)
	l.PanicIf(err)
}
