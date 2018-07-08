package handler

import (
	"fmt"
	"net/http"
	"payment/logger"
	"payment/order/services"
	utils "payment/utils/http"
)

func PaymentHandler(rw http.ResponseWriter, rq *http.Request) {
	parameters := utils.ExtractParametersFrom(rq, "account", "amount", "currency")
	account, amount, currency := parameters[0], parameters[1], parameters[2]
	logger.Logger.Printf("operation : account %s : %s %s\n", account, amount, currency)

	if services.CheckAccount(account) {
		rw.WriteHeader(201)
		fmt.Fprint(rw, "accepted\n")
		return
	}
	fmt.Fprint(rw, "rejected\n")
}
