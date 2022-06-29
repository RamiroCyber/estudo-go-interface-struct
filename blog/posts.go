package blog

import "github.com/ramiro/generics/blog/entities"

func GetPosts() []entities.Post {
	return []entities.Post{
		{
			ID: 1, Category: &entities.Category{ID: 1, Name: "Generics"},
			Title: "Aprendendo Generics",
			Text:  "Aqui vai o blog sobre generics",
			Slug:  "aprendendo-generics",
		},
		{
			ID: 2, Category: &entities.Category{ID: 2, Name: "Tutoriais"},
			Title: "Tutorial sobre Go",
			Text:  "Aqui vai o blog sobre Go",
			Slug:  "tutorial-sobre-go",
		},
		{
			ID: 3, Category: &entities.Category{ID: 3, Name: "Videos "},
			Title: "Aprendendo sobre ediçao de videos",
			Text:  "Aqui vai o blog sobre ediçao de videos",
			Slug:  "aprendendo-sobre-edicoes-de-videos",
		},
	}

}
