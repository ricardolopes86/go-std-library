package main

import (
	"log"
	"os"
)

type messageType int

const (
	INFO messageType = 0 + iota
	WARN
	ERR
	FATAL
)

func main(){
	writeLog(WARN, "our message in the logs")
	writeLog(INFO, "simple information")
	writeLog(ERR, "printing an error to the logs")
	writeLog(FATAL, "after this, the program should die!")
}

func writeLog(messageType messageType, message string) {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	switch messageType {
	case INFO:
		logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)
	case WARN:
		logger := log.New(file, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)
	case ERR:
		logger := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)
	case FATAL:
		logger := log.New(file, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Fatal(message)
	}
}