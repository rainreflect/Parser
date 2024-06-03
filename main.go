package main

import (
	_ "github.com/lib/pq"

	"github.com/rainreflect/parser/classificator"
	"github.com/rainreflect/parser/parser"
)

func main() {
	urls := []string{"https://www.freecodecamp.org/news/design-patterns-for-distributed-systems",
		"https://cyberleninka.ru/article/n/balansirovka-nagruzki-na-prilozhenie-ot-infrastruktury-do-bazy-dannyh",
		"https://theappsolutions.com/blog/development/high-load-systems/"}

	items, err := parser.ParseData(urls)
	if err != nil {
		panic(err)
	}
	// connStr := "host=localhost port=5432 user=postgres1 password=postgres1 dbname=labascraper sslmode=disable"
	// db, err := sql.Open("postgres", connStr)
	// if err != nil {
	// 	log.Fatalf("Error connecting to the database: %s", err)
	// }
	// defer db.Close()

	// err = myDB.CreateTable(db)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, item := range items {
	// 	myDB.SaveItemToDB(db, item)
	// }

	// tmpl, err := template.ParseFiles("template.html")
	// if err != nil {
	// 	log.Fatalf("Ошибка загрузки шаблона: %s", err)
	// }

	// http.HandleFunc("/articles", handler.ArticlesHandler(items, tmpl))

	// // Запуск HTTP-сервера
	// log.Println("Сервер запущен на http://localhost:8080")
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	log.Fatalf("Ошибка запуска сервера: %s", err)
	// }
	for _, item := range items {
		classificator.Classificate(item.Text)
	}
}
