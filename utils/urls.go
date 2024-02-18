package utils

import (
	"fmt"
	"sblog/config"
	m "sblog/models"

	"github.com/a-h/templ"
)

func CategoryUrl(category m.Category, page uint) templ.SafeURL {
	return templ.URL(fmt.Sprintf("%scategories/%s/%d.html", config.UrlBase, category.Slug, page))
}

func PostUrl(post m.Post) templ.SafeURL {
	return templ.URL(fmt.Sprintf("%spost/%d.html", config.UrlBase, post.Id))
}
