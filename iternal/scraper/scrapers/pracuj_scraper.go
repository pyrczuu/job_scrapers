package scrapers

import (
	"context"
	"github.com/gocolly/colly"
	"github.com/pfczx/jobscraper/iternal/scraper"
	"time"
)

const (
	titleSelector       = "h1[data-scroll-id='job-title']"
	companySelector     = ""
	locationSelector    = ""
	salarySelector      = ""
	linkSelector        = ""
	descriptionSelector = ""
	skillsSelector      = ""
)

type PracujScraper struct {
	timeoutBetweenScraps time.Duration
	collector            *colly.Collector
}

func NewPracujScraper() *PracujScraper {
	c := colly.NewCollector(
		colly.AllowedDomains("www.it.pracuj.pl", "it.pracuj.pl"),
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*pracuj.pl*",
		Parallelism: 2,
		RandomDelay: 2 * time.Second,
	})

	return &PracujScraper{
		timeoutBetweenScraps: 10 * time.Second,
		collector:            c,
	}
}

func (*PracujScraper) Source() string {
	return "pracuj.pl"
}

func (*PracujScraper) Scrape(ctx context.Context, q chan<- scraper.JobOffer) error {

	return nil
}
