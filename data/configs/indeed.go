package configs

import (
	"log/slog"

	"github.com/Wacky404/lurchers/data"

	"github.com/gocolly/colly"
)

type CollyCfg struct {
	C    *colly.Collector
	Data *data.Job
}

func newCollyCfg() *CollyCfg {
	return &CollyCfg{
		C:    colly.NewCollector(colly.Async(true)),
		Data: data.NewJob(),
	}
}

func IndeedConfig() *CollyCfg {
	cfg := newCollyCfg()
	cfg.Data.Posting.Website = "https://indeed.com/"
	// testing this out; will need to build this
	cfg.Data.Posting.Url = "https://www.indeed.com/jobs?q=%2B&l=Little+Rock%2C+AR&fromage=7&salaryType=%2440%2C000%2B&radius=5&jlid=68f779f7b0e38e09&rbl=Little+Rock%2C+AR&from=searchOnDesktopSerp&vjk=7a14f77130c03202"
	cfg.C.OnHTML("a[id^='job_']", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		cfg.C.Visit(e.Request.AbsoluteURL(link))
	})
	cfg.C.OnHTML("h1[class^='jobserch-JobInfoHeader-title']", func(e *colly.HTMLElement) {
		jobPosition := e.Text
		slog.Info("Job Found", slog.String("Position", jobPosition))
	})

	return cfg
}
