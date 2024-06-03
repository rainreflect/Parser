package parser

import (
	"strings"

	"github.com/rainreflect/parser/domain"
	"github.com/rainreflect/parser/scraper"
)

func ParseData(urls []string) ([]domain.Items, error) {
	var items []domain.Items
	var parser scraper.Parser

	for _, url := range urls {
		switch {
		case strings.Contains(url, "freecodecamp"):
			parser = scraper.NewFCCParser()
		case strings.Contains(url, "cyberleninka"):
			parser = scraper.NewCLParser()
		case strings.Contains(url, "theappsolutions"):
			parser = scraper.NewASParser()
		}

		item, err := parser.Parse(url)
		if err != nil {
			return nil, err
		}
		item.URL = url
		items = append(items, item)
	}
	return items, nil
}
