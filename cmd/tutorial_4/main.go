package main

import "fmt"

func main() {
/* 
	//arrays
	var intArr [3]int32
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])
	
	//Memmory location Array
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])
	*/
	
	//var intArr [3]int32 = [3]int32{1, 2, 3}
	intArr := [...]int32{1, 2, 3}
	fmt.Println(intArr)
	
	// slices
	var intSlice []int32 = []int32{4, 5, 6}
	//fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	//fmt.Println(intSlice)
	fmt.Printf("\nThe length is %v with capacity %v", len(intSlice), cap(intSlice))
	
	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println("\n", intSlice)
	
	var intSlices3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlices3)
	
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)
	
	var myMap2 = map[string]uint8{"adam":23, "sarah":45}
	fmt.Println(myMap2["adam"])
	// value is not stored
	fmt.Println(myMap2["jason"])
	
	var age, ok = myMap2["adam"]
	//delete(myMap2, "adam")
	if ok {
		fmt.Printf("The age is %v", age)
	}else {
		fmt.Printf("Invalid Name")
	}
	
	// loops
	
	for name, age := range myMap2 {
		fmt.Printf("\nName: %v, Age:%v \n", name, age)
	}

	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v \n", i ,v)
	}

	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}
	//second one
	
	for i:=0; i<10; i++ {
		fmt.Println(i)
	}

}