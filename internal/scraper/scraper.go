package scraper

type Scraper interface {
	Scrape() (interface{}, error)
}
