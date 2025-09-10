package dataservice

import (
	"New/model"
	"database/sql"
)

func CreateBook(db *sql.DB, book model.Book) error {
	query := `INSERT INTO BOOKS(id, title, author, year) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, book.Id, book.Title, book.Author, book.Year)
	if err != nil {
		return err
	}
	return nil
}
