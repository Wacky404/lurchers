/*
* TODO: Pagination works for Crunchyroll store for manga-books but
* Price selector is not working. Need to fix. Then persisting storage
 */
package configs

import (
	"context"
	"log"
	"log/slog"
	"time"

	"github.com/Wacky404/lurchers/evade"
	"github.com/Wacky404/lurchers/util"
	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
)

func main() {
	logFile, err := util.SetupLogger(util.WithLogName("logs/lurchers.log"))
	if err != nil {
		log.Fatal("error setting up logger", err)
	}
	defer logFile.Close()

	// this time out value will change
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()

	err = godotenv.Load()
	if err != nil {
		slog.Error("error loading .env file", slog.Any("error", err))
	}

	// bookShelve := []data.Book{}
	c := colly.NewCollector()
	evade.NewUserAgent(ctx, c)
	proxies := []string{util.GetVar("TOR", "socks5://127.0.0.1:9050")}
	err = evade.RotateProxy(c, &proxies)
	if err != nil {
		slog.Error("error configuring the RotateProxy", slog.Any("error", err))
	}

	c.OnHTML("div[class]", func(e *colly.HTMLElement) {
		className := e.Attr("class")
		pricemsrp := e.Attr("content")
		if className == "product" {
			// Replace multiple newlines with a single space
			// cleaned := strings.ReplaceAll(e.DOM.Text(), "\n", " ")
			// remove multiple spaces
			// cleaned = strings.Join(strings.Fields(cleaned), " ")
			title := e.ChildText(".pdp-link")
			priceDiscout := e.ChildText(".sales")
			slog.Info("Product Found", slog.String("Title", title), slog.String("Price", pricemsrp), slog.String("Sale", priceDiscout))
		}
	})

	c.OnHTML("a.image-tile-container", func(e *colly.HTMLElement) {
		srcLink := e.ChildText("src")
		slog.Debug(srcLink)
	})

	c.OnHTML("div.pagination-block ul li a", func(e *colly.HTMLElement) {
		rightArrow := e.Attr("class")
		if rightArrow == "right-arrow" {
			link := e.Attr("href")
			c.Visit(e.Request.AbsoluteURL(link))
		}
	})

	// before making a request print "Visiting..."
	c.OnRequest(func(r *colly.Request) {
		slog.Info("Visiting", slog.String("Request URL", r.URL.String()))
	})

	// start scraping on crunchyroll
	c.Visit("https://store.crunchyroll.com/collections/manga-books/?srule=Most-Popular")
	c.Wait()
}
