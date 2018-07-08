package handler

import (
	"fmt"
	"net/http"
	"payment/check/services"
	"payment/logger"
	utils "payment/utils/http"
)

func PaymentHandler(rw http.ResponseWriter, rq *http.Request) {
	parameters := utils.ExtractParametersFrom(rq, "account", "amount", "currency")
	account, _, _ := parameters[0], parameters[1], parameters[2]
	logger.Logger.Printf("check : account %s\n", account)

	if services.CheckAccount(account) {
		fmt.Fprint(rw, "accepted\n")
		return
	}
	fmt.Fprint(rw, "rejected\n")
}
