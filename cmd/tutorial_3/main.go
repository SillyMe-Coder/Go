package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello World"
	printMe(printValue)
	
	var numerator int = 11
	var denomintor int = 0
	var result, reminder, err = intDivision(numerator, denomintor)
/* 	if err != nil {
		fmt.Printf(err.Error())
	}else if reminder == 0 {
		fmt.Printf("The result of the integer division is %v", result)
	}else {
		fmt.Printf("The result of the integer division is %v with reminder %v\n", result, reminder)
		} */
	//fmt.Println(result)
	
	//switch statement
	switch{ 
		case err != nil: 
			fmt.Printf(err.Error())
		case reminder == 0:
			fmt.Printf("The result of the integer division is %v", result)
		default:
			fmt.Printf("The result of the integer division is %v with reminder %v\n", result, reminder)
	}
	switch reminder {
		case 0:
			fmt.Printf("The division is was exact")
		case 1, 2:
			fmt.Printf("The Division was close")
		default:
			fmt.Printf("The Divison was not Close")
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error){
	var err error
	if denominator == 0 {
		err = errors.New("Cannot Divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var reminder int = numerator % denominator
	return result, reminder, err
}