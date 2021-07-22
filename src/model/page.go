package model

import (

)

type Page struct{
	Title string
	Body string			//ini berguna buat TF-IDF
	URL string
	LinksTo []string	//ini berguna buat pageRank Algorithm
	Relevancy float64	//ini untuk nyimpen nilai pageRank Algorithm
}

