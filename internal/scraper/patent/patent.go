package patent

import (
	"strings"

	"github.com/gocolly/colly"
	"github.com/rs/zerolog/log"
)

type Patent struct {
	scrapper *colly.Collector
	options  []PatentOptions
	event    chan map[string]interface{}
}

type PatentOptions struct{}

type PatentModuler string

const (
	Pharsight PatentModuler = "https://pharsight.greyb.com/drug-patent-expiration-lists"
	Elixir    PatentModuler = "https://elixirdemo.greyb.com/drug-screener"
)

func (p *Patent) Scrapping() error {
	// Scrap method
	return p.scrapingPharsight()
}

func (p *Patent) Name() string {
	return "patent"
}

func (p *Patent) scrapingPharsight() error {
	// Scraping method
	p.scrapper.OnRequest(func(r *colly.Request) {
		log.Info().Msgf("Visiting %s", r.URL.String())
	})
	p.scrapper.OnHTML(".el-link", func(e *colly.HTMLElement) {
		log.Info().Msgf("Link found: %q -> %s", e.Text, e.Attr("href"))

		_ = e.Request.Visit(e.Attr("href"))
	})

	p.scrapper.OnHTML("tr.patent-table", func(e *colly.HTMLElement) {
		resp := map[string]interface{}{}
		resp["source"] = "pharsight"
		e.ForEach("td[data-label]", func(i int, e *colly.HTMLElement) {
			cleanText := strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(e.DOM.Text(), "\n", ""), "\t", ""), "  ", " "))
			label := e.Attr("data-label")

			resp[label] = cleanText
		})
		p.sendEvent(resp)
	})

	return p.scrapper.Visit(string(Pharsight))
}

func (p *Patent) sendEvent(data map[string]interface{}) {
	p.event <- data
}

func (p *Patent) GetEvent(fn func(data map[string]interface{})) {
	go func() {
		for {
			select {
			case data := <-p.event:
				fn(data)
			}
		}
	}()
}

func (p *Patent) RegisterCollector(collector interface{}, opts ...interface{}) {
	if collector == nil {
		p.scrapper = colly.NewCollector()
	} else {
		p.scrapper = collector.(*colly.Collector)
	}
	for _, opt := range opts {
		if optObj, ok := opt.(PatentOptions); ok {
			p.options = append(p.options, optObj)
		}
	}
	p.event = make(chan map[string]interface{})
}
