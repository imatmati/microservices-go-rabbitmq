package language

import "account/logger"

func PanicIf(err error) {
	if err != nil {
		logger.Logger.Println(err.Error())
		panic(err.Error())
	}
}
