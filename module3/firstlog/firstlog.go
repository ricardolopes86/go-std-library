package main

import (
	"log"
	"os"
)

func main(){
	file, err := os.Open("log.txt")
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	log.Println("First log message!")
}