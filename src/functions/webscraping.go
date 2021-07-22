package functions

import (
	"encoding/json"
	"strings"
	"log"
	// "net/http"
	"io/ioutil"
	"github.com/gocolly/colly"

	"search-engine-extended/src/model"
)

func removeDuplicateURL(URLs []string) []string {
	keys := make(map[string]bool)
    retVal := []string{}
  
    // If the key(values of the slice) is not equal
    // to the already present value in new slice (retVal)
    // then we append it. else we jump on another element.
    for _, entry := range URLs {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            retVal = append(retVal, entry)
        }
    }
    return retVal
}

func DoWebScrape() (error) {
	c := colly.NewCollector(
		colly.AllowedDomains("en.tempo.co"),
		colly.MaxDepth(2),
	)

	// pageCollector := c.Clone()

	var pages []model.Page

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// If link start with prefix, start scaping the page under the link found
		if !strings.HasPrefix(link, "https://en.tempo.co/read/") {
			return
		}

		e.Request.Visit(link)
	})

	c.OnHTML(`main`, func(e *colly.HTMLElement) {
		log.Println("start scraping: ", e.Request.URL)

		if !(e.Request.URL.String()=="https://en.tempo.co/news") {				//page awal gausah discrape
			hyperlink := removeDuplicateURL(e.ChildAttrs("a.terkini","href"))
			page := model.Page{
				Title: 		e.ChildText("h1"),
				Body:		e.ChildText("div#isi>p"),
				URL: 		e.Request.URL.String(),
				LinksTo: 	hyperlink,
			}
			
			pages = append(pages, page)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})


	c.Visit("https://en.tempo.co/news")

	file,_ := json.MarshalIndent(pages,""," ")

	_ = ioutil.WriteFile("documents.json", file, 0644)
	
	return nil
}
