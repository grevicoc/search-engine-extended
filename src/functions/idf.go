package functions

import (

)

//menerima inputan slice of string yang berisi banyak dokumen dan inputan kata-kata query. Akan dicari idf tiap-tiap kata
//ASUMSI arrOfQuery telah "bersih" alias tidak ada stopwords
func Idf(documents []string, arrOfQuery []string) (map[string]int)  {

	//pertama-tama bikin slice of BoW dari masing-masing document dari documents
	var bagOfWords []map[string]int
	for _, doc := range documents{
		bagOfWord,_ := countWords(doc)
		bagOfWords = append(bagOfWords, bagOfWord)
	}

	//selanjutnya bikin retval
	var retVal map[string]int = map[string]int{}
	for _,word := range arrOfQuery{
		var count = 0
		for _, bow := range bagOfWords{
			if _, find := bow[word]; find {
				count++
			}
		}

		retVal[word] = count
	}

	return retVal
}