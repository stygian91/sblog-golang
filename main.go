package main

import (
	m "sblog/models"
	"sblog/render"
	"time"
)

func main() {
	post := m.Post{
		Id:        1,
		Title:     "lorem",
		Content:   "# Lorem\n## Ipsum **dolor**",
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

	err := render.RenderSingle(post, categories)
	if err != nil {
		panic(err)
	}
}
