package main

import "fmt"

type gasEngine struct {
	mpg uint8
	gallons uint8

}

func main() {

	var  myEngine gasEngine = gasEngine{30, 10}

	fmt.Println(myEngine.mpg, myEngine.gallons)
}