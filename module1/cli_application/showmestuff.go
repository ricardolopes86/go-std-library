package main

import (
	"runtime"
	"fmt"
)

func main(){
	// this is where stuff starts to happen
	fmt.Printf("Go current version is %v running in %v. Sweet!!\n", runtime.Version(), runtime.GOOS)
}