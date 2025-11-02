package scrapers

import (
	"context"
	"github.com/gocolly/colly"
	"github.com/pfczx/jobscraper/iternal/scraper"
	"time"
)

const (
	titleSelector         = "h1[data-scroll-id='job-title']"
	companySelector       = "h2[data-scroll-id='employer-name']"
	locationSelector      = "div[data-test='offer-badge-title']"
	descriptionSelector   = `ul[data-test="text-about-project"]`                                                         //concat in code
	skillsSelector        = `span[data-test="item-technologies-expected"], span[data-test="item-technologies-optional"]` //concat in code
	salarySectionSelector = `div[data-test="section-salaryPerContractType"]`
	salaryAmountSelector  = `div[data-test="text-earningAmount"]`
	contractTypeSelector  = `span[data-test="text-contractTypeName"]`
z)

type PracujScraper struct {
	timeoutBetweenScraps time.Duration
	collector            *colly.Collector
}

funcNewPracujScraper() *PracujScraper {
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

func (p *PracujScraper) Scrape(ctx context.Context, urls []string, q chan<- scraper.JobOffer) error {
	for _, url := range urls {
	  time.Sleep(p.timeoutBetweenScraps)
		p.collector.OnHTML("html", func(e *colly.HTMLElement) {
			select {
			case <-ctx.Done():
				return
			default:
			}

			var job scraper.JobOffer
			job.URL = url
			job.Source = p.Source()
			job.Title = e.ChildText(titleSelector)
			job.Company = e.ChildText(companySelector)
			job.Location = e.ChildText(locationSelector)

			// desc
			e.ForEach(descriptionSelector+" li", func(_ int, el *colly.HTMLElement) {
				if text := el.Text; text != "" {
					job.Description += text + "\n"
				}
			})

			// skills
			var skills []string
			e.ForEach(skillsSelector, func(_ int, el *colly.HTMLElement) {
				if text := el.Text; text != "" {
					skills = append(skills, text)
				}
			})
			job.Skills = skills

			// salary
			e.ForEach(salarySectionSelector, func(_ int, el *colly.HTMLElement) {
				amount := el.ChildText(salaryAmountSelector)
				ctype := el.ChildText(contractTypeSelector)
				switch ctype {
				case "umowa o pracę":
					job.SalaryEmployment = amount
				case "umowa zlecenie":
					job.SalaryContract = amount
				case "kontrakt B2B":
					job.SalaryB2B = amount
				}
			})

			select {
			case <-ctx.Done():
				return
			case q <- job:
			}
		})

		if err := p.collector.Visit(url); err != nil {
			return err
		}
	}

	p.collector.Wait()
	return nil
}
