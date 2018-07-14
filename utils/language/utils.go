package language

import (
	"account/logger"
	"strings"
)

func PanicIf(err error, msgs ...string) {
	if err != nil {
		msg := strings.Join(msgs, " : ")
		// Replace double quote by single quote for JSON sake.
		msg = strings.Replace(msg, "\"", "'", -1)
		errMsg := strings.Replace(err.Error(), "\"", "'", -1)
		separator := " | "
		if len(msgs) == 0 {
			separator = ""
		}
		logger.Logger.Panic(msg, separator, errMsg)
	}
}
