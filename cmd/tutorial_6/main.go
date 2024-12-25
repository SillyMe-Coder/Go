package main

import (
	"fmt"
	"structs"
)

type gasEngine struct {
	mpg uint8
	gallons uint8

}

func main() {

	var  myEngine gasEngine = gasEngine{30, 10}
	var myEngine2 = struct {
		mpg uint8
		gallons unit8
	} {25, 15}// not reusable

	fmt.Println(myEngine.mpg, myEngine.gallons)
}