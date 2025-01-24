package evade

import (
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

func NewUserAgent(c *colly.Collector) {
	extensions.RandomUserAgent(c)
}


