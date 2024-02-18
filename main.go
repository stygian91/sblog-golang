package main

import (
	"sblog/config"
	m "sblog/models"
	"sblog/render"
	"time"
)

func main() {
	if err := config.Init(); err != nil {
		panic(err)
	}

	post := m.Post{
		Id:        1,
		Title:     "lorem",
		Content:   "# Lorem\n## Ipsum **dolor**\nLorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet. Nisi anim cupidatat excepteur officia. Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident. Nostrud officia pariatur ut officia. Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat officia voluptate. Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia eiusmod. Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim. Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis.",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Categories: []m.Category{
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
		},
	}

	if err := render.RenderSingle(post); err != nil {
		panic(err)
	}

	if err := render.RenderListingPage([]m.Post{post, post}, 1); err != nil {
		panic(err)
	}
}
