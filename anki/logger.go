package anki

import (
	"log"
	"os"
)

type Logger interface {
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}

func DefaultLogger() Logger {
	return log.New(os.Stdout, "anki: ", log.LstdFlags)
}
