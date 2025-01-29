package scraper

// ├── scraper/
// │   │   ├── lawsuits.go               # Scraping logic for lawsuit data
// │   │   ├── patent.go                 # Scraping logic for patent data
// │   │   ├── recall.go                 # Scraping logic for product recall data
// │   │   ├── registration.go           # Scraping logic for registration data
// │   │   └── manufacturer.go           # Scraping logic for manufacturer details
type ScraperModuler interface {
	Scrapping() ([]byte, error)
}
