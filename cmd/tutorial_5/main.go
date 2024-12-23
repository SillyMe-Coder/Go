package main

import ("fmt"
		"strings"
)
func main() {

	var myString = []rune("résumé")
	var indexed = myString[1]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString{
		fmt.Println(i, v)
	}
	fmt.Printf("\nThe lenght of 'mystring is %v", len(myString))

	var myRune = 'a'
	fmt.Printf("\nmyRune = %v", myRune)

	var strSlice = []string{"I", "L", "O", "V", "E", "C", "O", "D", "E"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])	
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)


}