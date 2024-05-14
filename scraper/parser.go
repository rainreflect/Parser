package scraper

import "github.com/rainreflect/parser/domain"

type Parser interface {
	Parse(url string) (domain.Items, error)
}
