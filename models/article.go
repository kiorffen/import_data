package models

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"import_data/database"
)

type Article struct {
	ID          string `json:"id"`
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	URL         string `json:"url" binding:"required"`
	PublishDate string `json:"publish_date"`
	Summary     string `json:"summary"`
	Tags        string `json:"tags"`
	Author      string `json:"author"`
}

func GenerateMD5(url string) string {
	hash := md5.Sum([]byte(url))
	return fmt.Sprintf("%x", hash)
}

func (a *Article) Save() error {
	a.ID = GenerateMD5(a.URL)
	
	var exists bool
	checkQuery := "SELECT EXISTS(SELECT 1 FROM articles WHERE id = ?)"
	err := database.DB.QueryRow(checkQuery, a.ID).Scan(&exists)
	if err != nil {
		return err
	}
	
	if exists {
		return fmt.Errorf("article with URL already exists")
	}

	query := `INSERT INTO articles (id, title, content, url, publish_date, summary, tags, author) 
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	
	var publishDate sql.NullString
	if a.PublishDate != "" {
		publishDate.String = a.PublishDate
		publishDate.Valid = true
	}

	var summary sql.NullString
	if a.Summary != "" {
		summary.String = a.Summary
		summary.Valid = true
	}

	var tags sql.NullString
	if a.Tags != "" {
		tags.String = a.Tags
		tags.Valid = true
	}

	var author sql.NullString
	if a.Author != "" {
		author.String = a.Author
		author.Valid = true
	}

	_, err = database.DB.Exec(query, a.ID, a.Title, a.Content, a.URL, 
		publishDate, summary, tags, author)
	return err
}