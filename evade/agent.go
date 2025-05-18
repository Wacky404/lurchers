package evade

import (
    "context"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

func NewUserAgent(ctx context.Context, c *colly.Collector) {
	extensions.RandomUserAgent(c)
}


