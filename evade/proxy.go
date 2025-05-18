package evade

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
)

func RotateProxy(c *colly.Collector, proxies *[]string) error {
	rp, err := proxy.RoundRobinProxySwitcher(*proxies...)
	if err != nil {
		return fmt.Errorf("error setting up proxy switcher: %s", err.Error())
	}
	c.SetProxyFunc(rp)

	return nil
}
