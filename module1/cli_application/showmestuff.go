package main

import (
	"runtime"
	"fmt"
)

func main(){
	// this is where stuff starts to happen
	fmt.Println("Go current version is" + runtime.Version())
}