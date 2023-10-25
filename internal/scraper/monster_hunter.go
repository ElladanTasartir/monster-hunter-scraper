package scraper

import (
	"github.com/gocolly/colly"
)

type MonsterHunterScraper struct {
	collector *colly.Collector
	address   string
}

type Monster struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

type MonstersResponse struct {
	Monsters []Monster `json:"monsters"`
}

func NewMonsterHunterScraper(address string) (*MonsterHunterScraper, error) {
	collector := colly.NewCollector(
		colly.CacheDir("./"))

	return &MonsterHunterScraper{
		collector: collector,
		address:   address,
	}, nil
}

func (s *MonsterHunterScraper) Scrape() (interface{}, error) {
	var monsters []Monster

	s.collector.OnHTML("div#gallery-1 > .wikia-gallery-item", func(e *colly.HTMLElement) {
		monster := Monster{}

		e.ForEach(".thumb img", func(_ int, item *colly.HTMLElement) {
			monster.Image = item.Attr("data-src")
		})

		e.ForEach(".lightbox-caption > a", func(_ int, item *colly.HTMLElement) {
			monster.Name = item.Text
		})

		monsters = append(monsters, monster)
	})

	err := s.collector.Visit(s.address)
	if err != nil {
		return nil, err
	}

	return monsters, nil
}
