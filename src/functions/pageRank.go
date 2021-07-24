package functions

import (
	"fmt"
	"search-engine-extended/src/model"
)

//ini documents masih ga merubah
func PageRank(documents *[]model.Page) {
	//pertama dibikin map dengan key URL dan val objek Page
	var mapDocuments map[string]*model.Page = map[string]*model.Page{}
	for idx,doc := range *documents{
		mapDocuments[doc.URL] = &(*documents)[idx]
	}

	baseImportance := float32(1) / float32(len(*documents))
	for _,val := range mapDocuments{
		countLinks := float32(len(val.LinksTo))
		for _,link := range val.LinksTo{
			fmt.Println(baseImportance)
			mapDocuments[link].Importance += baseImportance / countLinks
		}
	}
}