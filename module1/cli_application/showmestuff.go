package main

import (
	"runtime"
	"fmt"
)

func main(){
	// this is where stuff starts to happen
	fmt.Printf("Go current version is %v Sweet!!\n", runtime.Version())
}