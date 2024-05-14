package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	_ "github.com/lib/pq"

	myDB "github.com/rainreflect/parser/db"
	"github.com/rainreflect/parser/domain"
	"github.com/rainreflect/parser/handler"
	"github.com/rainreflect/parser/scraper"
)

func main() {
	urls := []string{"https://www.freecodecamp.org/news/design-patterns-for-distributed-systems",
		"https://cyberleninka.ru/article/n/balansirovka-nagruzki-na-prilozhenie-ot-infrastruktury-do-bazy-dannyh",
		"https://theappsolutions.com/blog/development/high-load-systems/"}
	items, err := ParseData(urls)
	if err != nil {
		panic(err)
	}
	connStr := "host=localhost port=5432 user=postgres1 password=postgres1 dbname=labascraper sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}
	defer db.Close()

	err = myDB.CreateTable(db)
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range items {
		myDB.SaveItemToDB(db, item)
	}

	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		log.Fatalf("Ошибка загрузки шаблона: %s", err)
	}

	http.HandleFunc("/articles", handler.ArticlesHandler(items, tmpl))

	// Запуск HTTP-сервера
	log.Println("Сервер запущен на http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Ошибка запуска сервера: %s", err)
	}

}

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
