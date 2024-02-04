package main

import (
	"context"
	"os"
	"sblog/config"
	"sblog/templates"
	m "sblog/models"
	"time"
)

func main() {
	err := config.Init()
	if err != nil {
		panic(err)
	}

	p := m.Post{
		Title:     "lorem",
		Content:   "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet. Nisi anim cupidatat excepteur officia. Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident. Nostrud officia pariatur ut officia. Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat officia voluptate. Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia eiusmod. Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim. Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis.",
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
