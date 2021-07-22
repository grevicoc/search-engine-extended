package functions

import (
	// "encoding/json"
	"fmt"
	// "log"
	// "net/http"
	"github.com/gocolly/colly"

	"search-engine-extended/src/model"
)

func DoWebScrape() ([]model.Page, error) {
	c := colly.NewCollector()

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.Visit("https://en.tempo.co/news")
	
	return nil, nil
}