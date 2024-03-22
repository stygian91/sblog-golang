package render

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sblog/models"
	views "sblog/templates"

	"github.com/a-h/templ"
)

func prepareOutSubdir(path string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	dir := filepath.Join(cwd, "out/" + path)
	if err = os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	return dir, nil
}

func renderToFile(dest string, component templ.Component) error {
	fileHandle, err := os.Create(dest)
	if err != nil {
		return err
	}

	return component.Render(context.Background(), fileHandle)
}

func RenderSingle(post models.Post) error {
	dir, err := prepareOutSubdir("post")
	if err != nil {
		return err
	}

	dest := filepath.Join(dir, fmt.Sprintf("%d.html", post.Id))
	return renderToFile(dest, views.PostSingle(post))
}

func RenderListingPage(posts []models.Post, page uint64, maxPages uint64) error {
	dir, err := prepareOutSubdir("page")
	if err != nil {
		return err
	}

	dest := filepath.Join(dir, fmt.Sprintf("%d.html", page))
	return renderToFile(dest, views.PostListing(posts, page, maxPages))
}
