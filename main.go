package main

import (
	"fmt"
	"github.com/ramiro/generics/blog"
	"github.com/ramiro/generics/blog/entities"
	"github.com/ramiro/generics/cache"
)

func main() {
	cc := cache.New[entities.Category]()
	categories := blog.GetCategories()
	for _, value := range categories {
		cc.Set(value.Slug, value)
	}

	cp := cache.New[entities.Post]()
	posts := blog.GetPosts()
	for _, value := range posts {
		cp.Set(value.Slug, value)
	}

	fmt.Println(cc.Get("generics"))
	fmt.Println("--------------------")
	fmt.Println(cp.Get("aprendendo-generics"))
}
