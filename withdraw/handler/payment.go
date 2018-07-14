package handler

import (
	"account/logger"
	utils "account/utils/http"
	l "account/utils/language"
	"account/withdraw/services"
	"fmt"
	"net/http"
	"strconv"
)

//WithdrawHandler handles withdraw operation on account.
func WithdrawHandler(rw http.ResponseWriter, rq *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			rw.WriteHeader(500)
			rw.Write([]byte(utils.ErrorJSON(fmt.Sprintf("%v", r))))
		}
	}()
	// I - Parameters extraction
	parameters := utils.ExtractParametersFrom(rq, "account", "amount", "currency")
	account, currency := parameters[0], parameters[2]
	amount, err := strconv.Atoi(parameters[1])
	l.PanicIf(err, "Conversion error of amount")
	logger.Logger.Printf("operation : account %s : %d/100 of %s\n", account, amount, currency)

	statusCode := 500
	result := "rejected"

	// II - Check account
	if services.CheckAccount(account) {
		logger.Logger.Printf("Account %s passed check\n", account)
		// III - Update amount of account
		amount = services.UpdateAccount(account, amount, currency)
		logger.Logger.Printf("Account %s updated\n", account)
		statusCode = 201
		result = "accepted"
	}
	// IV - Send response
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	fmt.Fprint(rw, fmt.Sprintf("{\"account\":\"%s\",\"result\":\"%s\"}", account, result))

}
