package main

import (
	"fmt"
	"log"

	"github.com/kevincobain2000/go-glasssdoor-scraper/surf"
)

func main() {
	// set env and flags
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	sg := surf.NewScrapeGlassdoor()
	j := sg.ParseFlag().Scrape().ToJSON()

	fmt.Println(j)
}
