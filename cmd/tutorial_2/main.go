package main

import "fmt"
import "unicode/utf8"

func main() {
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)
	
	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1/intNum2)
	fmt.Println(intNum1%intNum2)
	
	var myString string = "Hello" +" "+ "Me"
	fmt.Println(myString)
	
	fmt.Println(utf8.RuneCountInString("γ"))
	
	var myRune rune = 'a'
	fmt.Println(myRune)
	
	var myBoolean bool = false
	fmt.Println(myBoolean)
	
	//SHort hand
	myVar := "Hello, Me"
	fmt.Println(myVar)
	
	//multi vars
	var1, var2 := 1, 2
	fmt.Println(var1, var2)
	
	//import some function
	//var myVar1 string = foo()
	//fmt.Println(myVar1)
	
	//constants
	const myConst string = "const value"
	fmt.Println(myConst)
	
	const pi float32 = 3.1415
}
