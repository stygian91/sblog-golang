package main

import (
	"context"
	"os"
	"sblog/config"
	m "sblog/models"
	"sblog/templates"
	"time"
)

func main() {
	err := config.Init()
	if err != nil {
		panic(err)
	}

	p := m.Post{
		Title:     "lorem",
		Content:   "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	categories := []m.Category{
		{
			Id:   1,
			Name: "Ipsum",
			Slug: "ipsum",
		},
		{
			Id:   2,
			Name: "Dolor",
			Slug: "dolor",
		},
	}

	comp := templates.PostSingle(p, categories)
	comp.Render(context.Background(), os.Stdout)
}
