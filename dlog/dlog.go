package dlog

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
)

var (
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errLogger   *log.Logger
	debugLogger *log.Logger
	printStdout bool
	errorsOnly  bool
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

	infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	warnLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime)
	errLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime)
	debugLogger = log.New(file, "DEBUG: ", log.Ldate|log.Ltime)
}

func PrintToStdout(errorsOnlyToStdout bool) {
	printStdout = true
	errorsOnly = errorsOnlyToStdout
}

func PrintOnlyToFile() {
	printStdout = false
}

func formatMessage(msg string, params ...interface{}) string {
	_, fileName, lineNumber, ok := runtime.Caller(2)
	if ok {
		formattedMsg := fmt.Sprintf(msg, params...)
		formattedMsg = fmt.Sprintf("%s:%d: %s\n", path.Base(fileName), lineNumber, formattedMsg)
		return formattedMsg
	} else {
		return fmt.Sprintf(msg+"\n", params...)
	}
}

func Info(msg string, params ...interface{}) {
	formattedMsg := formatMessage(msg, params...)

	if printStdout && errorsOnly {
		fmt.Printf(formattedMsg)
	}

	infoLogger.Printf(formattedMsg)
}

func Warn(msg string, params ...interface{}) {
	formattedMsg := formatMessage(msg, params...)

	if printStdout && errorsOnly {
		fmt.Printf(formattedMsg)
	}

	warnLogger.Printf(formattedMsg)
}

func Err(msg string, params ...interface{}) {
	formattedMsg := formatMessage(msg, params...)

	if printStdout && errorsOnly {
		fmt.Printf(formattedMsg)
	}

	errLogger.Printf(formattedMsg)
}

func Debug(msg string, params ...interface{}) {
	formattedMsg := formatMessage(msg, params...)
	if printStdout && errorsOnly {
		fmt.Printf(formattedMsg)
	}

	debugLogger.Printf(formattedMsg)
}
