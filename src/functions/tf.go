package functions

import (

)

func countWords(doc string) (map[string]int, error){
	//inisialisasi retval
	var retVal map[string]int = map[string]int{}

	//menyiapkan tempWord untuk variabel sementara sebuah kata dan iterasi tiap char di document
	var tempWord string = ""
	for _, char := range doc{
		if char!=' '{
			tempWord = tempWord + string(char)	
		}

		if char==' '{
			retVal[tempWord]++
			tempWord = ""
		}
	}

	return retVal, nil
}

func divideByTotal(bagOfWord map[string]int) (map[string]float32, error){
	//hitung total kata
	var totalWords int = 0
	for _, val:= range bagOfWord{
		totalWords+= val
	}
	
	//inisialisasi retval
	var retVal map[string]float32 = map[string]float32{}

	for key, val:= range bagOfWord{
		retVal[key] = float32(val) / float32(totalWords)
	}

	return retVal, nil
}

func Tf(doc string) (map[string]float32, error){
	bow,_ := countWords(doc)
	result,_ := divideByTotal(bow)

	return result, nil
}
