package scraper

import (
	"github.com/noydhiet/mandrill-scrapper/internal/scraper/patent"
)

// ├── scraper/
// │   │   ├── lawsuits.go               # Scraping logic for lawsuit data
// │   │   ├── patent.go                 # Scraping logic for patent data
// │   │   ├── recall.go                 # Scraping logic for product recall data
// │   │   ├── registration.go           # Scraping logic for registration data
// │   │   └── manufacturer.go           # Scraping logic for manufacturer details
type ScraperModuler interface {
	Scrapping() error
	RegisterCollector(collector interface{}, opts ...interface{})
	GetEvent(fn func(data map[string]interface{}))
	Name() string
}

type Scraper struct {
	ScraperModuler
}

func NewScraper(scraper ScraperModuler) *Scraper {
	return &Scraper{
		scraper,
	}
}

func NewPatentScraper(collector interface{}, opts ...interface{}) *Scraper {
	obj := NewScraper(&patent.Patent{})
	obj.RegisterCollector(collector, opts)
	return obj
}
