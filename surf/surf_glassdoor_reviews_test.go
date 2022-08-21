package surf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScrape(t *testing.T) {
	sg := NewScrapeGlassdoor()
	reviewURL := "https://www.glassdoor.com.au/Reviews/GitHub-Reviews-E671945_P19.htm?filter.iso3Language=eng"
	sg.ReviewsURL = &reviewURL
	sg.Scrape()
	assert.Greater(t, len(sg.Pros), 0)
	assert.Greater(t, len(sg.Cons), 0)
	assert.Greater(t, len(sg.Authors), 0)
	assert.Greater(t, len(sg.Ratings), 0)
}
