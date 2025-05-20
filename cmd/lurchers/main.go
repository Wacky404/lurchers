package main

import (
	"context"
	"log"
	"log/slog"
	"time"

	"github.com/Wacky404/lurchers/data/configs"
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

	// our buffed collector for indeed
	i := configs.IndeedConfig()
	evade.NewUserAgent(ctx, i.C)
	proxies := []string{util.GetVar("TOR", "socks5://127.0.0.1:9050")}
	err = evade.RotateProxy(i.C, &proxies)
	if err != nil {
		slog.Error("error configuring the RotateProxy", slog.Any("error", err))
	}

	// before making a request print "Visiting..."
	i.C.OnRequest(func(r *colly.Request) {
		slog.Info("Going to website", slog.String("Request URL", r.URL.String()))
	})
	// start scraping on website(s)
	i.C.Visit(i.Data.Posting.Url)
	i.C.Wait()
}
