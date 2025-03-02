package data

import "strings"

type Query struct {
	Title        string `json:"title"`
	TotalResults string `json:"totalResults"`
	SearchTerms  string `json:"searchTerms"`
	Count        int    `json:"count"`
	StartIndex   int    `json:"startIndex"`
	Safe         string `json:"safe"`
}

type Queries struct {
	PreviousPage []Query `json:"previousPage"`
	Request      []Query `json:"request"`
	NextPage     []Query `json:"nextPage"`
}

// Point to Note, if the title starts with Openings At or Careers At, it is a company advertising their jobs
type SearchResult struct {
	Title        string `json:"title"`
	Link         string `json:"link"`
	FormattedUrl string `json:"formattedUrl"`
}

func (result *SearchResult) IsValidJobListing() bool {
	resultTitle := strings.ToLower(result.Title)
	keyWordsThatIndicateListing := []string{"opening", "careers"}
	for _, keyword := range keyWordsThatIndicateListing {
		if strings.Contains(resultTitle, keyword) {
			return true
		}
	}
	return false
}
