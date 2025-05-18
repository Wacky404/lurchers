package configs

import (
	"github.com/Wacky404/lurchers/data"
	"github.com/gocolly/colly"
)

func indeedConfig() *colly.Collector {
	job := new(data.Job)
	c := colly.NewCollector()
	c.OnHTML("a[id^='job_']", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		c.Visit(e.Request.AbsoluteURL(link))
	})
	c.OnHTML("h1[class^='jobserch-JobInfoHeader-title']", func(e *colly.HTMLElement) {

	})

	return c
}
