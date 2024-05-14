package scraper

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rainreflect/parser/domain"
)

type AppSolutionsParser struct {
}

func NewASParser() *AppSolutionsParser {
	return &AppSolutionsParser{}
}

func (p AppSolutionsParser) Parse(url string) (domain.Items, error) {
	var items domain.Items
	res, err := http.Get(url)
	if err != nil {
		log.Printf("error fetching URL %s: %v", url, err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Printf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Printf("error reading document %s: %v", url, err)
	}

	headline := doc.Find("h1").Text()
	items.Title = headline

	author := doc.Find("div.post-author__name").Text()

	items.Author = author
	var textBuilder strings.Builder

	doc.Find("h2#conclusion").Each(func(i int, s *goquery.Selection) {

		for sibling := s.Next(); sibling.Length() > 0; sibling = sibling.Next() {
			textBuilder.WriteString(sibling.Text() + "\n")
		}
	})

	items.Summary = textBuilder.String()
	doc.Find("div.post-content").Each(func(i int, s *goquery.Selection) {
		items.Text = s.Text()
	})

	return items, nil
}
