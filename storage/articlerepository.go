package storage

import (
	"fmt"
	"log"

	"github.com/evgborovoy/StandardWebServer/internal/app/models"
)

type ArticleRepository struct {
	storage *Storage
}

var (
	tableArticle = "articles"
)

// Добавить статью в БД
func (ar *ArticleRepository) Create(article *models.Article) (*models.Article, error) {
	query := fmt.Sprintf("INSERT INTO %s (title, author, content) VALUES($1, $2, $3) RETURNING id", tableArticle)
	if err := ar.storage.db.QueryRow(
		query,
		article.Title,
		article.Author,
		article.Content).Scan(&article.ID); err != nil {
		return nil, err
	}
	return article, nil
}

// удалить статью по id
func (ar *ArticleRepository) DeleteByID(id int) (*models.Article, error) {
	article, ok, err := ar.FindByID(id)
	if err != nil {
		return nil, err
	}
	if ok {
		query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", tableArticle)
		_, err = ar.storage.db.Exec(query, id)
		if err != nil {
			return nil, err
		}
	}
	return article, nil
}

// Найти статью по id
func (ar *ArticleRepository) FindByID(id int) (*models.Article, bool, error) {
	founded := false
	articles, err := ar.SelectAll()
	if err != nil {
		return nil, founded, err
	}
	var articleFinded *models.Article
	for _, ar := range articles {
		if ar.ID == id {
			articleFinded = ar
			founded = true
			break
		}
	}
	return articleFinded, founded, nil
}

func (ar *ArticleRepository) SelectAll() ([]*models.Article, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableArticle)
	rows, err := ar.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	articles := make([]*models.Article, 0)
	for rows.Next() {
		ar := models.Article{}
		err := rows.Scan(&ar.ID, &ar.Title, &ar.Author, &ar.Content)
		if err != nil {
			log.Println(err)
			continue
		}
		articles = append(articles, &ar)
	}
	return articles, nil
}
