package main

import (
	"context"
	"log"
	"log/slog"
	"time"

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

	// before making a request print "Visiting..."
	c.OnRequest(func(r *colly.Request) {
		slog.Info("Visiting", slog.String("Request URL", r.URL.String()))
	})
	// start scraping on website(s)
	c.Visit("https://store.crunchyroll.com/collections/manga-books/?srule=Most-Popular")
	c.Wait()
}
