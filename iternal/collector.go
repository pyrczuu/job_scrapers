package iternal

import (
	"context"
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/pfczx/jobscraper/database"
	"github.com/pfczx/jobscraper/iternal/scraper"
)

func StartCollector(ctx context.Context, db *sql.DB,scrapers []scraper.Scraper){
	out :=  scraper.RunScrapers(ctx,scrapers)
	querier := database.New(db)
	go func(){
		for job := range out{
			log.Printf("Saving job: %s from %s",job.Title,job.Company)
			params := database.CreateJobOfferParams{
				ID: uuid.New(),
				Title: job.Title,
				Company: job.Company,
				Location: job.Location,
				Description: job.Description,
				Url: job.URL,
				
			}
			if _,err := querier.CreateJobOffer(ctx,params) ; err !=nil{
				log.Printf("Error in saving: %s from %s",job.Title,job.Company)
			}
		}
	}()
}

