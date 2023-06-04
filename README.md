# Go Crawler

A simple internal link crawler written in Go.

## Description

This bot will scrape the DOM for all `<a>` tags, follow the URL, and log a line to file with the URL and HTTP status code.

## Getting Started
### Dependencies

* [gocolly/colly](https://github.com/gocolly/colly) scraping library
* Zero dependancy executable (thanks Go!)

### Installing
* Download the executable for your platform (MacOS, Windows, Linux)

### Executing program
#### Mac or Linux
* Execute the file with the flag `-domain` to crawl a website (minus the https://)
```bash
./crawler -domain example.com
```

#### Windows
* Execute the file with the flag `-domain` to crawl a website (minus the https://)
```bash
./crawler.exe -domain example.com
```

## Help
* Feel free to raise an issue or submit a pull request.