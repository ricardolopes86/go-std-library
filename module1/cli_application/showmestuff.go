package main

import (
	"runtime"
	"fmt"
	"os"
	"bufio"
)

func main(){
	// this is where stuff starts to happen
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Whats your name?")
	text, _ := reader.ReadString('\n')
	fmt.Printf("Hello %v", text)
	fmt.Printf("Go current version is %v running in %v. Sweet!!\n", runtime.Version(), runtime.GOOS)
}