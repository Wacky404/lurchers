package main

import (
	"github.com/Wacky404/lurchers/data"
	"github.com/gocolly/colly"
)

func main() {
	bookShelve := []data.Book{}
	c := colly.NewCollector()
	c.OnHTML("div[class]", func(e *colly.HTMLElement) {
		className := e.Attr("class")
		if className == "product" {
			b := new(data.Book)
		}
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		c.Visit(e.Request.AbsoluteURL(link))
	})
}
