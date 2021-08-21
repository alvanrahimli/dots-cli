package dlog

import (
	"log"
	"os"
	"path"
	"strings"
)

var (
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errLogger   *log.Logger
	debugLogger *log.Logger
)

func init() {
	homeDir, homeErr := os.UserHomeDir()
	if homeErr != nil {
		return
	}

	logFilePath := path.Join(homeDir, ".dots-cli.log")
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile|log.LstdFlags)
	warnLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	errLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger = log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	infoLogger.Printf("========| APP STARTED (command: %s) |========\n", strings.Join(os.Args, " "))
}

func Info(msg string, params ...interface{}) {
	infoLogger.Printf(msg+"\n", params...)
}

func Warn(msg string, params ...interface{}) {
	warnLogger.Printf(msg+"\n", params...)
}

func Err(msg string, params ...interface{}) {
	errLogger.Printf(msg+"\n", params...)
}

func Debug(msg string, params ...interface{}) {
	debugLogger.Printf(msg+"\n", params...)
}
