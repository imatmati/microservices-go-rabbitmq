package handler

import (
	"account/check/services"
	"account/logger"
	utils "account/utils/http"
	"fmt"
	"net/http"
)

func CheckHandler(rw http.ResponseWriter, rq *http.Request) {
	parameters := utils.ExtractParametersFrom(rq, "account", "amount", "currency")
	account, _, _ := parameters[0], parameters[1], parameters[2]
	logger.Logger.Printf("check : account %s\n", account)
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(rw, "{\"account\": \"%s\", \"check\": %t}", account, services.CheckAccount(account))
}
