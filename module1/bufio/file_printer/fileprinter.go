package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "--help" {
		fmt.Println("Usage: fileprinter filename.(extension)")
	} else {
		if len(args) == 0 {
			fmt.Println("Please, provide file name.")
		} else {
			fmt.Println("Please, choose an option to print the file:")
			fmt.Println("1: UPPER CASE")
			fmt.Println("2: Title Caps")
			fmt.Println("3: lower case")
			var option int
			fmt.Scanf("%d", &option)
			
			switch option {
			case 1:
				file, err := os.Open(args[0])
				if err != nil {
					fmt.Println(err)
				}
				defer file.Close()

				scanner := bufio.NewScanner(file)
				for scanner.Scan(){
					fmt.Println(strings.ToUpper(scanner.Text()))
				}

				if err := scanner.Err(); err != nil {
					fmt.Println(err)
				}
			case 2:
				file, err := os.Open(args[0])
				if err != nil {
					fmt.Println(err)
				}
				defer file.Close()

				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					fmt.Println(strings.ToTitle(scanner.Text()))
				}

				if err := scanner.Err(); err != nil {
					fmt.Println(err)
				}
			case 3:
				file, err := os.Open(args[0])
				if err != nil {
					fmt.Println(err)
				}

				defer file.Close()

				scanner := bufio.NewScanner(file)
				for scanner.Scan(){
					fmt.Println(strings.ToLower(scanner.Text()))
				}

				if err := scanner.Err(); err != nil {
					fmt.Println(err)
				}
			default:
				fmt.Println("Wrong option.")
			}
		}
	}
}