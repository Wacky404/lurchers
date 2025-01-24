package evade

import (
	"errors"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
)

func RotateProxy(c *colly.Collector, proxies *[]string) error {
	rp, err := proxy.RoundRobinProxySwitcher(*proxies...)
	if err != nil {
		return errors.New("proxies were not set for collector")
	}
	c.SetProxyFunc(rp)

	return nil
}
