package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	// "log"
	"os"
	"search-engine-extended/src/functions"
	"search-engine-extended/src/model"
)

//TODO integrasiin heapsort sama model
//TODO FRONTEND WEBSITE ASQYIAJDIDHIUWHDU
func main(){
	var testArray = []float32{0.27,11.1,5.4,6.02,3.01,3.01}
	functions.HeapSort(testArray)
	fmt.Println(testArray)

	//ambil documents dari file
	jsonFile, err := os.Open("documents.json")
	if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Successfully Opened documents.json")
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var documents []model.Page
	json.Unmarshal(byteValue, &documents)

	// fmt.Println(documents)

	//calculate tf-idf value and cosine-sim then save it on struct val
	query := map[string]int{"jokowi":1,"covid":1,"19":1}
	var valueIDF = functions.Idf(documents,query)

	for idx,doc := range documents{
		valueTF,_ := functions.Tf(doc)
		tfidf := functions.TfIdf(valueTF,valueIDF)
		doc.Relevancy = functions.CosineSim(query,tfidf)
		// fmt.Println(valueTF)
		// fmt.Println(valueIDF)
		// fmt.Println(tfidf)
		fmt.Println(doc.Relevancy)
		// fmt.Println(" ")
		documents[idx] = doc
	}

	functions.PageRank(&documents)
	

	for idx,doc := range documents{
		fmt.Printf("Document %d, Relevancy: %f | Importance: %f\n", idx, doc.Relevancy, doc.Importance)
	}

	/* WEBSCRAPING, UNCOMMENT UNTUK MELAKUKAN WEBSCRAPE*/
	// err := functions.DoWebScrape()
	// if err!=nil{
	// 	fmt.Println(err)
	// }
}
