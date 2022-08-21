<p align="center">
  <a href="https://github.com/kevincobain2000/go-glasssdoor-scraper">
    <img alt="go-glasssdoor-scraper" src="logo.png" width="360">
  </a>
</p>

<h3 align="center">Scrape Glassdoor Reviews, Ratings, Author, and Timestamps.</h3>

<p align="center">
  Works as headless browser
  <br>
  Scrapes Glassdoor reviews for a given employer
  <br>
  No Glassdoor account required
  <br>
  Written in Go :heart:

</p>

**Supports:** Glassdoor reviews scrape paginated URLS.

**Command line:** Arch free binary to run as scheduler on any platform.

**Headless:** Uses headless browser to run without selenium or chromium drivers.

**Proxy support:** Works behind a proxy.

**Dependencies:** None.


## Installation

```sh
go install github.com/kevincobain2000/go-glasssdoor-scraper@latest
```

### From Command Line:

```sh
go-glasssdoor-scraper -reviews-url=https://www.glassdoor.com.au/Reviews/GitHub-Reviews-E671945.htm
```

```json
{
    "ReviewsURL": "https://www.glassdoor.com.au/Reviews/GitHub-Reviews-E671945_P19.htm?filter.iso3Language=eng",
    "EmployerName": "Github",
    "PaginatedURLS": [
        "https://www.glassdoor.com.au/Reviews/GitHub-Reviews-E671945_P19.htm?filter.iso3Language=eng",
        "https://glassdoor.com.au/Reviews/GitHub-Reviews-E671945_P20.htm?filter.iso3Language=eng"
    ],
    "Pros": [
        "The best part about working at GitHub is the great colleagues! They are uniformly friendly, super talented, and have a positive, supportive attitude. All the engineers I have worked with have been really motivated and produce high-quality work.\n\nA majority of GitHub employees work remotely; some in small satellite offices but most at home/cafes/shared office space/whatever. Because of this, the company goes to a lot of effort to make remote work effective. Communications are largely asynchronous (e.g., via GitHub issues and pull requests), most meetings are streamed and recorded, and there are lots of chances to meet other GitHubbers when necessary: the annual company summit, mini-summits for departments or teams kicking off a project, plus we are welcome to travel to visit HQ or other GitHubbers when we feel the urge to connect. Most \"synchronous\" discussion takes place in online chat, though even here transcripts are kept for anybody to read. Time zones can be a real hassle but otherwise it all works quite well. It's so awesome being able to work for a San Francisco-based company while being able to live almost anywhere!\n\nThe company is very transparent. Many company-wide issues are discussed in the open, and anybody's input is welcome.\n\nUntil the beginning of 2014, GitHub didn't have any formal management structure to speak of. The result was predictable: people worked on what they found exciting, but it was hard to gather critical mass to implement large-scale projects. Since then people-managers have been introduced, which has improved the company focus. So far management has used a relatively light touch. In particular, work is still mostly organized by the engineers working on it rather than top-down by management. In my opinion the introduction of management has improved the company, though it has definitely caused a shift in the culture.\n\nGitHub the company also seems to have a good heart. As far as I can tell, everybody at the company from top to bottom is really trying hard to do the right thing:\n\n* There is a constant internal dialog about how to improve diversity, both within GitHub (e.g., by doing everything we can to eliminate bias in our recruiting processes) and also within the whole industry (for example by supporting organizations and events that are targeted at underrepresented groups). This isn't just an HR goal; we are all thinking about it.\n* We do everything we can to support open-source software. Most obviously, we provide free Git hosting to open-source projects, and we give our free users every bit the same level of support as we give to our paying customers. We contribute internally-developed code to upstream open-source projects as much as possible. We have \"open-source Fridays\" twice a month when GitHubbers are encouraged to work on open source projects. We sponsor many open-source conferences and other events.\n* GitHub promotes a healthy work-life balance. I guess (being alpha-geeks) a lot of us work too many hours. But that pressure doesn't come from the company. I've never seen anything but encouragement for people announcing vacation (which is unlimited) or taking parental leave (4 months PTO for new parents).\n\nFinally, it's great to build a product that millions of people use and love. It's a pleasure to run into GitHub users in real life because they are so positive about the company. We're helping to propagate the open-source ethos:\n\n* We provide a great place for open-source projects to host their code.\n* Companies use GitHub to promote code-sharing and communal development internally, and often to open-source internal projects.\n* There are even governments that put draft legislation in GitHub and allow citizens to propose pull requests to change it.\n\nReading back over what I've written, it sounds like I'm a marketing shill. But in fact I'm just a very happy software engineer. After nearly two years at the company I am still almost giddy that I get to work at GitHub.",
        "I get paid extremely well to work with caring people on challenging engineering problems for a product that I believe in from the comfort of my home office. Oh, and that product is GitHub! This is the best job I've ever had.",
        "Management is open to new ideas, quick growth, opportunity to have a fingerprint in company.",
        ...
    ],
    "Cons": [
        "Remote work impedes serendipitous communication. It's an effort to keep up with what is happening in the rest of the company when you don't bump into your colleagues in the kitchen or in the elevator.\n\nAs a hangover from the management-free days, it can sometimes still be hard to figure out who is in charge of what.",
        "GitHub is growing rapidly, so there's a lot to do. It's easy to overcommit yourself into a bad work/life balance, or find yourself in need of work from an overloaded team. As long as you're mindful of those edges and ensure you're working on important things, you can create quite a place for yourself here.",
        ...
    ],
    "Authors": [
        "Jun 8, 2015 - Software Engineer in Jakarta",
        "Jul 18, 2015 - Anonymous Employee in Jakarta",
        ...
    ],
    "Ratings": [
        "5.0",
        "5.0",
        ...
    ]
}
```

--

### Usage in GO

```go
import (
    "github.com/kevincobain2000/go-glasssdoor-scraper/surf"
)
func main() {
	sg := surf.NewScrapeGlassdoor()
	reviewURL := "https://www.glassdoor.com.au/Reviews/GitHub-Reviews-E671945.htm"
	sg.ReviewsURL = &reviewURL
	sg.Scrape()

    fmt.Print(sg.Pros)
    fmt.Print(sg.Cons)
    fmt.Print(sg.Authors)
    fmt.Print(sg.Ratings)
}
```

### CHANGE LOG

- v1.0 - Initial release includes Glassdoor reviews scraper

### ROADMAP

- v1.1 - Output JSON to file
