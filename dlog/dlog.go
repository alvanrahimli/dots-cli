package dlog

import (
	"log"
	"os"
	"time"
)

var (
	InfoLogger  *log.Logger
	WarnLogger  *log.Logger
	ErrLogger   *log.Logger
	DebugLogger *log.Logger
)

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarnLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	DebugLogger = log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	InfoLogger.Printf("========| APP STARTED at %s |========\n",
		time.Now().Format("02/01/2006 15:04:05"))
}

func Info(msg string, params ...interface{}) {
	newLine := msg + "\n"
	InfoLogger.Printf(newLine, params...)
}

func Warn(msg string, params ...interface{}) {
	newLine := msg + "\n"
	WarnLogger.Printf(newLine, params...)
}

func Err(msg string, params ...interface{}) {
	newLine := msg + "\n"
	ErrLogger.Printf(newLine, params...)
}

func Debug(msg string, params ...interface{}) {
	newLine := msg + "\n"
	DebugLogger.Printf(newLine, params...)
}
