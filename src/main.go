package main

import (
	"fmt"
	"search-engine-extended/src/functions"
)

func main(){
	var testArray = []float32{0.27,11.1,5.4,6.02,3.01,3.01}
	functions.HeapSort(testArray)
	fmt.Println(testArray)

	testDoc := "Lorem ipsum dolor amet ahay de kocak ipsum lorem test aja ini mah ahay ahay de"
		
	value,_ := functions.Tf(testDoc)
	fmt.Println(value)
}