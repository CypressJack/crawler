package main

import (
	"flag"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	// Define flags
	domain := flag.String("domain", "", "Domain to crawl (don't add https://)")
	flag.Parse()

	// Check if required flag exists
	if !flag.Parsed() || *domain == "" {
		fmt.Println("Domain is required, remember to omit https:// ")
		flag.Usage()
		os.Exit(1)
	}

	if strings.Contains(*domain, "https://") || strings.Contains(*domain, "http://") {
		fmt.Println("Please omit https:// from the domain, eg. google.com")
		os.Exit(1)
	}

	// Set up logging
	logfile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logfile.Close()

	log.SetOutput(logfile)

	// Set up crawler
	c := colly.NewCollector(
		colly.AllowedDomains(*domain),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob: *domain + "/*",
		// Set a delay between requests to these domains
		Delay: 1 * time.Second,
		// Add an additional random delay
		RandomDelay: 1 * time.Second,
	})

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println(r.StatusCode, r.Request.URL, err)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println(r.StatusCode, r.Request.URL)
	})

	c.Visit("https://" + *domain)
}
