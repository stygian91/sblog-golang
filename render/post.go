package render

import (
	"context"
	"os"
	"path/filepath"
	m "sblog/models"
	"strconv"
	t "sblog/templates"
)

func RenderSingle(post m.Post, categories []m.Category) error {
	content, err := post.ContentToHtml()
	if err != nil {
		return err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	dir := filepath.Join(cwd, "out/post")
	dest := filepath.Join(dir, strconv.FormatUint(uint64(post.Id), 10) + ".html")

	if err = os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	fileHandle, err := os.Create(dest)
	if err != nil {
		return err
	}

	return t.PostSingle(post, content, categories).Render(context.Background(), fileHandle)
}
