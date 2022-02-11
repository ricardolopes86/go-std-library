package main

import (
	"fmt"
)

func main(){
	var name string
	fmt.Println("What is your name")
	fmt.Scanf("%s", &name) //use %q if you want to have more than 1 work input, just remember it should be in quotes
	fmt.Printf("Hello %s! Nice to meet you!\n", name)
}