package functions

import (
	// "fmt"
	"math"
	"search-engine-extended/src/model"
	"strings"
)

//fungsi antara untuk nentuin apakah suatu char ada di suatu kumpulan char
func contains(char rune, word []rune) (bool){
	for _,c := range word{
		if char==c {
			return true
		}
	}
	return false
}

//fungsi untuk menghitung suatu dokumen ada kata apa saja dan jumlahnya berapa
func countWords(doc string) (map[string]int, error){
	//inisialisasi retval
	var retVal map[string]int = map[string]int{}

	//inisialisasi huruf yang bisa menjadi "pengganggu" di sebuah kata
	var tandaBaca = []rune{'.','>',',','<','/','?','\'','"',';',':','\\','|',']','}','[','{','=','+','-','_',')','(','*','&','&','^','%','$','#','@','!','`','~',' ','\n'}

	//menyiapkan tempWord untuk variabel sementara sebuah kata dan iterasi tiap char di document
	var tempWord string = ""
	for _, char := range doc{

		//sebelum mengupdate, pastikan kata-katanya udh di-trim dengan dilowercase dan dihilangkan tanda bacanya
		if !contains(char,tandaBaca) {
			tempWord = tempWord + string(char)	
		}

		if contains(char,tandaBaca)  {
			tempWord = strings.ToLower(tempWord)
			retVal[tempWord]++
			tempWord = ""
		}
	}

	return retVal, nil
}

func countLogOfFreq(bagOfWord map[string]int) (map[string]float32, error){
	//hitung total kata
	var totalWords int = 0
	for _, val:= range bagOfWord{
		totalWords+= val
	}
	
	//inisialisasi retval
	var retVal map[string]float32 = map[string]float32{}

	for key, val:= range bagOfWord{
		// retVal[key] = float32(val) / float32(totalWords)
		retVal[key] = float32(math.Log(float64(1 + val)))
	}

	return retVal, nil
}

func Tf(doc model.Page) (map[string]float32, error){
	bow,_ := countWords(doc.Body)
	result,_ := countLogOfFreq(bow)

	return result, nil
}

//menerima inputan slice of string yang berisi banyak dokumen dan inputan kata-kata query. Akan dicari idf tiap-tiap kata
//ASUMSI arrOfQuery telah "bersih" alias tidak ada stopwords
func Idf(documents []model.Page, arrOfQuery map[string]int) (map[string]float32)  {

	//pertama-tama bikin slice of BoW dari masing-masing document dari documents
	var bagOfWords []map[string]int
	for _, doc := range documents{
		bagOfWord,_ := countWords(doc.Body)
		bagOfWords = append(bagOfWords, bagOfWord)
	}

	//selanjutnya bikin retval
	var retVal map[string]float32 = map[string]float32{}
	for key := range arrOfQuery {		var count = 0
		for _, bow := range bagOfWords{
			//ngecek apakah di dalam bow tersebut ada word yg dimaksud
			if bow[key]!=0{
				count++
			}
		}
		retVal[key] = float32(math.Log(float64(len(documents)) / float64(count)))
		
	}

	return retVal
}

func TfIdf(valueTF map[string]float32, valueIDF map[string]float32) (map[string]float32) {
	var retVal map[string]float32 = map[string]float32{}

	for key,val := range valueIDF{
		if valueTF[key]!=0.0{
			retVal[key] = valueTF[key] * val
		}else{
			retVal[key] = 0.0
		}
	}

	return retVal
}

//harusnya di sini vektor query sama vectorDocument isinya dah sama
func CosineSim(query map[string]int, vectorDocument map[string]float32) (float64){
	var dotProduct float32 = 0
	for key,val := range query{
		dotProduct += vectorDocument[key] * float32(val)
	}
	// fmt.Println(dotProduct)

	var doubleLineQuery float32 = 0
	for _,val := range query{
		doubleLineQuery += float32(val*val)
	}
	doubleLineQuery = float32(math.Sqrt(float64(doubleLineQuery)))
	// fmt.Println(doubleLineQuery)

	var doubleLineDocument float32 = 0
	for _,val := range vectorDocument{
		doubleLineDocument += val*val
	}
	doubleLineDocument = float32(math.Sqrt(float64(doubleLineDocument)))
	// fmt.Println(doubleLineDocument)

	if doubleLineDocument==0.0 {
		return 0
	}
	
	doubleLine := doubleLineQuery*doubleLineDocument
	// fmt.Println(dotProduct)
	// fmt.Println(doubleLine)

	return float64(dotProduct) / float64(doubleLine)
}