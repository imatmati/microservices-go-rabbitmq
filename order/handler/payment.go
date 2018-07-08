package handler

import (
	"fmt"
	"net/http"
	"payment/logger"
	"payment/order/services"
	utils "payment/utils/http"
	l "payment/utils/language"
	"strconv"
)

func PaymentHandler(rw http.ResponseWriter, rq *http.Request) {
	parameters := utils.ExtractParametersFrom(rq, "account", "amount", "currency")
	account, currency := parameters[0], parameters[2]
	amount, err := strconv.Atoi(parameters[1])
	l.PanicIf(err)
	logger.Logger.Printf("operation : account %s : %d %s\n", account, amount, currency)

	if services.CheckAccount(account) {
		logger.Logger.Printf("Checked\n")
		if err := services.UpdateAccount(account, amount, currency); err == nil {
			logger.Logger.Printf("Updated\n")
			rw.WriteHeader(201)
			fmt.Fprint(rw, "{'account':'"+account+"','result':'accepted'}")
			return
		}
	}
	rw.WriteHeader(200)
	fmt.Fprint(rw, "{'account':'"+account+"','result':'rejected'}")
}
