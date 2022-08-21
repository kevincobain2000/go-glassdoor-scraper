package surf

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/gosimple/slug"
)

// ScrapeGlassdoor for dui
type ScrapeGlassdoor struct {
	ReviewsURL    *string
	EmployerName  string
	PaginatedURLS []string
	Pros          []string
	Cons          []string
	Authors       []string
	Ratings       []string
}

// NewScrapeGlassdoor returns a new ScrapeGlassdoor instance
func NewScrapeGlassdoor() *ScrapeGlassdoor {
	return &ScrapeGlassdoor{}
}

func (sg *ScrapeGlassdoor) ParseFlag() *ScrapeGlassdoor {
	sg.ReviewsURL = flag.String("reviews-url", "", "Glassdoor's reviews URL: https://www.glassdoor.com.au/Reviews/GitHub-Reviews-E671945.htm")

	err := flag.CommandLine.Parse(os.Args[1:])

	if err != nil || *sg.ReviewsURL == "" {
		log.Fatal("[fatal] -reviews-url is required")
	}

	return sg
}

func (sg *ScrapeGlassdoor) ToJSON() string {
	log.Println("Output JSON")
	val, err := json.MarshalIndent(sg, "", "    ")
	if err != nil {
		return "{}"
	}
	return string(val)
}

func (sg *ScrapeGlassdoor) Scrape() *ScrapeGlassdoor {

	sg.setPaginatedURLS()
	for _, url := range sg.PaginatedURLS {
		sg.SetProsAndCons(url)
		//sleep for a while
		log.Println("[info] Sleeping crawl reviews to prevent 429")
		time.Sleep(time.Second * 1)
	}
	return sg
}

func (sg *ScrapeGlassdoor) setPaginatedURLS() {
	sg.PaginatedURLS = append(sg.PaginatedURLS, *sg.ReviewsURL)
	sg.recursiveCrawl(*sg.ReviewsURL)
}

func (sg *ScrapeGlassdoor) recursiveCrawl(url string) {
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{url},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {

			r.HTMLDoc.Find("link[rel='next']").Each(func(_ int, s *goquery.Selection) {
				nextURL, found := s.Attr("href")
				if found {
					sg.PaginatedURLS = append(sg.PaginatedURLS, nextURL)
					log.Print("[info] Next URL: ", nextURL)
					log.Println("[info] Sleeping recursive call to prevent 429")
					time.Sleep(time.Second * 1)
					sg.recursiveCrawl(nextURL)
				}
			})
		},
		RobotsTxtDisabled: true,
	}).Start()
}

func (sg *ScrapeGlassdoor) SetProsAndCons(url string) {
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{url},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			var pros []string
			var cons []string
			var authors []string
			var ratings []string
			var employerName string

			r.HTMLDoc.Find("span[data-test='pros']").Each(func(_ int, s *goquery.Selection) {
				pros = append(pros, s.Text())
			})
			r.HTMLDoc.Find("span[data-test='cons']").Each(func(_ int, s *goquery.Selection) {
				cons = append(cons, s.Text())
			})
			r.HTMLDoc.Find("span[class='authorInfo']").Each(func(_ int, s *goquery.Selection) {
				authors = append(authors, s.Text())
			})
			r.HTMLDoc.Find("span[class^='ratingNumber']").Each(func(_ int, s *goquery.Selection) {
				ratings = append(ratings, s.Text())
			})
			r.HTMLDoc.Find("p[class^='employerName']").Each(func(_ int, s *goquery.Selection) {
				employerName = slug.Make(s.Text())
			})

			if len(pros) != len(cons) {
				// case should not happen
				log.Println("[warn] pros, cons length mismatch")
				return
			}
			// empty authors
			if len(pros) != len(authors) {
				// create slice of empty strings
				for i := 0; i < len(pros); i++ {
					authors = append(authors, "")
				}
			}
			if len(pros) != len(ratings) {
				// create slice of empty strings
				for i := 0; i < len(pros); i++ {
					ratings = append(ratings, "")
				}
			}
			sg.Pros = append(sg.Pros, pros...)
			sg.Cons = append(sg.Cons, cons...)
			sg.Authors = append(sg.Authors, authors...)
			sg.Ratings = append(sg.Ratings, ratings...)
			if sg.EmployerName != "" {
				sg.EmployerName = employerName
			}
		},
		RobotsTxtDisabled: true,
	}).Start()
}
