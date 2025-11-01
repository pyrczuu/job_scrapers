package scraper

import (
	"context"
)

type JobOffer struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Company     string   `json:"company"`
	Location    string   `json:"location"`
	Salary      string   `json:"salary"`
	Description string   `json:"description"`
	URL         string   `json:"url"`
	Source      string   `json:"source"`
	PublishedAt *string  `json:"published_at,omitempty"` //potencial problems
	Skills      []string `json:"skills,omitempty"`
}

type Scraper interface {
	Source() string
	Scrape(ctx context.Context, q chan<- JobOffer) error
}
