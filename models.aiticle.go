package main

import (
	"errors"
)

type article struct {
	ID		int		`json:"id"`
	Title	string	`json:"title"`
	Content	string	`json:"content"`
}

var articleList = []article {
	article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
	article{ID: 3, Title: "Article 3", Content: "Article 3 body"},
	article{ID: 4, Title: "Article 4", Content: "Article 4 body"},
	article{ID: 5, Title: "Article 5", Content: "Article 5 body"},
}

func getAllArticles() []article {
	return articleList
}

func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("article not found")
}