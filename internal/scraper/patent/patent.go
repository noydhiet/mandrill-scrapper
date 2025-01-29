package patent

import (
	"github.com/gocolly/colly"
)

type Patent struct {
	apis     []string
	scrapper *colly.Collector
	options  map[string]string
}

var titleRefTargets = map[string]bool{
	"Recently expired drug patents":   true,
	"Recently added drug patents":     true,
	"Recently published drug patents": true,
	"Recent drug patent litigations":  true,
}

func (p *Patent) Scrapping() ([]byte, error) {
	// Scrap method

	p.scrapper.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if e.Text == "Recently expired drug patents" {

		}
		e.Request.Visit(e.Attr("href"))
	})

	return nil, nil
}
