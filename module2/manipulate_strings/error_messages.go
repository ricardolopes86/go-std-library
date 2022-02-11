package main

import (
	"bufio"
	"fmt"
	"os"
)

type messageType int

const (
	INFO messageType = 0 + iota
	WARN
	ERR
)

const (
	Infocolor = "\033[1;34m%s\033[0m"
	Warncolor = "\033[1;33m%s\033[0m"
	Errcolor = "\033[1;31m%s\033[0m"
)
func main(){
	filename := "test.txt"
	showMessage(INFO, fmt.Sprintf("Opening file %s", filename))
	showMessage(WARN, fmt.Sprint("This operation may fail if file doesn't exist\n"))

	file, err := os.Open(filename)
	if err != nil {
		showMessage(ERR, err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

func showMessage(messageType messageType, message string) {
	switch messageType {
	case INFO:
		printMessage := fmt.Sprintf("\nInfo: \n%s\n", message)
		fmt.Printf(Infocolor, printMessage)
	case WARN:
		printMessage := fmt.Sprintf("\nWarn: \n%s\n", message)
		fmt.Printf(Warncolor, printMessage)
	case ERR:
		printMessage := fmt.Sprintf("\nErr: \n%s\n", message)
		fmt.Printf(Errcolor, printMessage)
	}
}