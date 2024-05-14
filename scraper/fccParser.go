package scraper

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rainreflect/parser/domain"
)

type FreeCodeCampParser struct {
}

func NewFCCParser() *FreeCodeCampParser {
	return &FreeCodeCampParser{}
}

func (p FreeCodeCampParser) Parse(url string) (domain.Items, error) {
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
	doc.Find("h1").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		title = strings.TrimSpace(title)
		items.Title = title
	})
	doc.Find(".author-card-name").Each(func(i int, s *goquery.Selection) {
		author := s.Text()
		author = strings.TrimSpace(author)
		items.Author = author
	})
	doc.Find(".summary").Each(func(i int, s *goquery.Selection) {
		items.Summary = s.Text()
	})
	doc.Find(".post-content ").Each(func(i int, s *goquery.Selection) {
		items.Text = s.Text()
	})

	return items, nil
}
