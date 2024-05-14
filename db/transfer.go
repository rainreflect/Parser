package myDB

import (
	"database/sql"
	"log"

	"github.com/rainreflect/parser/domain"
)

func CreateTable(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS myschema.articles (
        id SERIAL PRIMARY KEY,
        title TEXT,
        author TEXT,
        summary TEXT,
        text TEXT,
        url TEXT
    )`
	_, err := db.Exec(query)
	return err
}

func SaveItemToDB(db *sql.DB, item domain.Items) error {
	query := `INSERT INTO myschema.articles (title, author, summary, text, url) VALUES ($1, $2, $3, $4, $5)`
	result, err := db.Exec(query, item.Title, item.Author, item.Summary, item.Text, item.URL)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	log.Printf("Rows affected: %d\n", rowsAffected)
	return nil
}

func GetArticles(db *sql.DB) ([]domain.Items, error) {
	rows, err := db.Query(`SELECT id, title, author, summary, text, url FROM myschema.articles`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []domain.Items
	for rows.Next() {
		var item domain.Items
		if err := rows.Scan(&item.Title, &item.Author, &item.Summary, &item.Text, &item.URL); err != nil {
			return nil, err
		}
		articles = append(articles, item)
	}
	return articles, nil
}
