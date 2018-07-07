package logger

import (
	"io"
	"log"
	"os"
)

var Logger *log.Logger

func InitLogger(output, prefix string) {
	var out io.Writer
	var err error
	switch {
	case output == "stdout":
		out = os.Stdout
	case output == "stderr":
		out = os.Stderr
	default:
		out, err = os.OpenFile(output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err.Error())
		}
	}
	Logger = log.New(out, prefix, log.Lmicroseconds)
}
