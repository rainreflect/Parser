package scraper

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rainreflect/parser/domain"
)

type CyberLeninkaParser struct {
}

func NewCLParser() *CyberLeninkaParser {
	return &CyberLeninkaParser{}
}

func (p CyberLeninkaParser) Parse(url string) (domain.Items, error) {
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
	headline := doc.Find("h1 i[itemprop='headline']").Text()
	items.Title = headline

	authorNotFixed := doc.Find("h2").Text()
	authorArr := strings.Split(authorNotFixed, " ")[12:14]
	author := strings.Join(authorArr, " ")

	items.Author = author

	summary := doc.Find("p[itemprop='description'] ").Text()

	items.Summary = summary
	doc.Find("div.ocr").Each(func(i int, s *goquery.Selection) {
		items.Text = s.Text()
	})
	return items, nil
}
