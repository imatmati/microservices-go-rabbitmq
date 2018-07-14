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

func WithdrawHandler(rw http.ResponseWriter, rq *http.Request) {
	parameters := utils.ExtractParametersFrom(rq, "account", "amount", "currency")
	account, currency := parameters[0], parameters[2]
	amount, err := strconv.Atoi(parameters[1])
	l.PanicIf(err)
	logger.Logger.Printf("operation : account %s : %d %s\n", account, amount, currency)

	if services.CheckAccount(account) {
		logger.Logger.Printf("Checked\n")
		if amount, err = services.UpdateAccount(account, amount, currency); err == nil {
			logger.Logger.Printf("%s pdated\n", account)
			rw.WriteHeader(201)
			fmt.Fprint(rw, "{\"account\":\""+account+"\",\"result\":\"accepted\",\"amount\":\""+strconv.Itoa(amount)+"\"}")
			return
		} else {
			logger.Logger.Println(err.Error())
		}
	}
	rw.WriteHeader(200)
	fmt.Fprint(rw, "{\"account\":\""+account+"\",\"result\":\"rejected\",\"error\":\""+err.Error()+"\"}")
}
