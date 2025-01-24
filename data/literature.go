package data

import (
	"image"
	"time"
)

type Author struct {
	name        string
	description string
	photo       image.Image
}

type Category struct {
	name        string
	subCategory SubCategory
}

type SubCategory struct {
	name string
}

type Publisher struct {
	name string
	logo image.Image
}

type Book struct {
	title                  string
	author                 Author
	category               Category
	series                 string
	description            string
	cover                  image.Image
	publisher              Publisher
	year_published         uint16
	print_length           uint16
	language               string
	isbn_10                uint32
	isbn_13                uint32
	reviews                string
	rating                 float32
	hcPriceUS              map[string]uint16
	hcPriceCAN             map[string]uint16
	pbPriceUS              map[string]uint16
	pbPriceCAN             map[string]uint16
	ebPriceUS              map[string]uint16
	ebPriceCAN             map[string]uint16
	audioPriceUS           uint16
	audioPriceCAN          uint16
	priceLastModified      time.Time
	cheapestPrice          map[string]uint16
	cheapPriceLastModified time.Time
}
