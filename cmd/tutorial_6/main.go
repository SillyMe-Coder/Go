package main

import "fmt"

type gasEngine struct {
	mpg uint8
	gallons uint8
	owner
}

type owner struct {
	name string
}

func main() {

	var  myEngine gasEngine = gasEngine{30, 10, owner{"John Doe"}}
	myEngine.mpg = 35
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name)
}