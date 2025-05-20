package data

import (
	"image"
	"time"
)

type Author struct {
	Name        string
	Description string
	Photo       image.Image
}

type Category struct {
	Name        string
	SubCategory SubCategory
}

type SubCategory struct {
	Name string
}

type Publisher struct {
	Name string
	Logo image.Image
}

type Book struct {
	Title                  string
	Author                 Author
	Category               Category
	Series                 string
	Description            string
	Cover                  image.Image
	Publisher              Publisher
	Year_published         uint16
	Print_length           uint16
	Language               string
	Isbn_10                uint32
	Isbn_13                uint32
	Reviews                string
	Rating                 float32
	HcPriceUS              map[string]uint16
	HcPriceCAN             map[string]uint16
	PbPriceUS              map[string]uint16
	PbPriceCAN             map[string]uint16
	EbPriceUS              map[string]uint16
	EbPriceCAN             map[string]uint16
	AudioPriceUS           uint16
	AudioPriceCAN          uint16
	PriceLastModified      time.Time
	CheapestPrice          map[string]uint16
	CheapPriceLastModified time.Time
}
