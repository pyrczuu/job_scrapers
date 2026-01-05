package main

import (
	"context"
	"database/sql"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pfczx/jobscraper/iternal"
	"github.com/pfczx/jobscraper/iternal/scraper"
	"github.com/pfczx/jobscraper/iternal/scraper/scrapers"
	//"github.com/pyrczuu/urlScraper"
	//"github.com/pyrczuu/nofluff_scraper"
	"github.com/pfczx/jobscraper/urlgoscraper"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	db, err := sql.Open("sqlite3", "./database/jobs.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var wg sync.WaitGroup

	//read only for js backend
	//_, err = db.Exec("PRAGMA journal_mode=WAL;")

	ctx := context.Background()

	var (
		noflufUrls,justjoinUrls []string
	)
	wg.Add(2)
/*
	go func() {
		defer wg.Done()
		pracujUrls = urlsgocraper.CollectPracujPl(ctx)
	}()
*/
	go func() {
		defer wg.Done()
		noflufUrls, _ = urlsgocraper.NofluffScrollAndRead(ctx)
	}()

	go func() {
		defer wg.Done()
		justjoinUrls, _ = urlsgocraper.JustJoinScrollAndRead(ctx)
	}()

	wg.Wait()


	scrapersList := []scraper.Scraper{
		scrapers.NewNoFluffScraper(noflufUrls),
		//scrapers.NewPracujScraper(pracujUrls),
		scrapers.NewJustJoinItScraper(justjoinUrls),
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		iternal.StartCollector(ctx, db, scrapersList)
	}()

	wg.Wait()
	log.Println("Scraping Completed")
}
