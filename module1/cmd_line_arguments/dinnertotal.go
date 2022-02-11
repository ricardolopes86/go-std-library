package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "--help" {
		fmt.Println("Usage: dinnertotal <Bill Total Amount> <Tip Percentage>")
		fmt.Println("Example: dinnertotal 30 10")
	} else {
		if len(args) != 2 {
			fmt.Println("You must enter 2 arguments.")
		} else {
			mealTotal, _ := strconv.ParseFloat(args[0], 32)
			tipAmount, _ := strconv.ParseFloat(args[1], 32)
			totalAmount := calcTotal(float32(mealTotal), float32(tipAmount))
			fmt.Printf("You meal total will be %.2f\n", totalAmount)
		}

	}
}

func calcTotal(mealTotal float32, tipAmount float32) float32{
	totalPrice := mealTotal + (mealTotal * (tipAmount /100))
	return totalPrice
}