package handler

import (
	"account/check/services"
	"account/logger"
	utils "account/utils/http"
	"fmt"
	"net/http"
)

func CheckHandler(rw http.ResponseWriter, rq *http.Request) {
	account := utils.ExtractParametersFrom(rq, "account")[0]
	logger.Logger.Printf("check : account %s\n", account)
	rw.Header().Set("Content-Type", "application/json")
	if check, err := services.CheckAccount(account); err == nil {
		fmt.Fprintf(rw, "{\"account\": \"%s\", \"check\": %t}", account, check)
	} else {
		fmt.Fprintf(rw, "{\"account\": \"%s\", \"check\": %t, \"error\":\"%s\"}", account, check, err.Error())
	}

}
