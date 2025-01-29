package lawsuit

import (
	"github.com/gocolly/colly"
)

type Lawsuit struct {
	// Lawsuit struct
	apis     []string
	scrapper *colly.Collector
}

func (l *Lawsuit) Scrapping() ([]byte, error) {
	// Scrap method
	return nil, nil
}
