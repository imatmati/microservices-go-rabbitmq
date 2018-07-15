package handler

import (
	"account/check/data"
	"account/check/services"
	"account/logger"
	utils "account/utils/http"
	l "account/utils/language"
	"fmt"
	"net/http"
	"strconv"
)

func CheckHandler(rw http.ResponseWriter, rq *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			rw.WriteHeader(500)
			//To avoid problem with runtime.errString when nil pointer error : fmt.Sprintf("%v", r)
			rw.Write([]byte(utils.ErrorJSON(fmt.Sprintf("%v", r))))
		}
	}()
	account := utils.ExtractParametersFrom(rq, "account")[0]
	logger.Logger.Printf("check : account %s\n", account)
	rw.Header().Set("Content-Type", "application/json")
	if checked, err := services.CheckAccount(account); err == nil {

		result := data.Result{Account: account, Check: strconv.FormatBool(checked)}
		if resultJson, err := result.ToJSON(); err == nil {
			fmt.Fprintf(rw, resultJson)
		} else {
			l.PanicIf(err, "Json marshalling error")
		}

	} else {
		l.PanicIf(err)
	}

}
