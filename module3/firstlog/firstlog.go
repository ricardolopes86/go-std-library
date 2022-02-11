package main

import (
	"log"
	"os"
)

type messageType int

var (
	WarningLogger 	*log.Logger
	InfoLogger 		*log.Logger
	ErrorLogger		*log.Logger
	FatalLogger		*log.Logger
)

func init() {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	InfoLogger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger := log.New(file, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	FatalLogger := log.New(file, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)

}

func main(){
	WarningLogger.Println("our message in the logs")
	InfoLogger.Println("simple information")
	ErrorLogger.Println("printing an error to the logs")
	FatalLogger.Println("after this, the program should die!")
}