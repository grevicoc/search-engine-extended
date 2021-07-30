package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	
	"net/http"
	"html/template"
	
	"search-engine-extended/src/functions"
	"search-engine-extended/src/model"
)

type M map[string]interface{}

func main(){
	/* Processing Documents */
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
	var documents []model.Page = []model.Page{}
	json.Unmarshal(byteValue, &documents)
	functions.PageRank(&documents)
	/* End Processing Documents */

	/* WEB */
	// ini tuh jadi setiap ada request ke /static/ bakal ngeakses . relatif dari folder assets di server
	http.Handle("/static/",
        http.StripPrefix("/static/",
            http.FileServer(http.Dir("assets"))))

	// ini buat ngerender semua template di dalem views biar siap dipake
	// var tmpl, err = template.ParseGlob("views/*")
    // if err != nil {
    //     panic(err.Error())
    //     return
    // }

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := M{}
		var tmpl = template.Must(template.ParseFiles(
			"views/index.html",
			"views/_page.html",
		))

		var err = tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
    })

	http.HandleFunc("/result", func(w http.ResponseWriter, r *http.Request) {
		if r.Method=="GET"{
			/* Ambil Value Form */
			if err := r.ParseForm(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			var rawQuery = r.FormValue("searchbar")
			// fmt.Println(rawQuery)
			query, _ := functions.CountWords(rawQuery)
			// fmt.Println(query)

			/* Calculate TF-IDF, Cosine Sim, and Pagerank */
			var valueIDF = functions.Idf(documents,query)

			for idx,doc := range documents{
				valueTF,_ := functions.Tf(doc)
				tfidf := functions.TfIdf(valueTF,valueIDF)
				doc.Relevancy = functions.CosineSim(query,tfidf)
				documents[idx] = doc
			}

			functions.HeapSortAdapted(documents)
			/* End Calculate */

			// for idx,doc := range documents{
			// 	fmt.Printf("Document %d, Relevancy: %f | Importance: %f\n", idx, doc.Relevancy, doc.Importance)
			// }

			/* Showing Result on Website */
			data := M{"results" : documents}
			var tmpl = template.Must(template.ParseFiles(
				"views/index.html",
				"views/_page.html",
			))

			err = tmpl.ExecuteTemplate(w, "index", data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
    })

    fmt.Println("server started at localhost:9000")
    http.ListenAndServe(":9000", nil)

	/* WEBSCRAPING, UNCOMMENT UNTUK MELAKUKAN WEBSCRAPE*/
	// err := functions.DoWebScrape()
	// if err!=nil{
	// 	fmt.Println(err)
	// }
}