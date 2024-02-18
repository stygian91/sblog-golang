package models

import (
	"sblog/config"
	"sblog/utils/markdown"
	"time"
	"unicode"
)

type Post struct {
	Id         uint
	Title      string
	Content    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Categories []Category
}

func (this Post) ContentToHtml() (string, error) {
	return markdown.ToString(this.Content)
}

func (this Post) Excerpt() string {
	wordCount := 0
	excerpt := ""
	isPrevSpace := false

	for _, ch := range this.Content {
		if !unicode.IsSpace(ch) {
			isPrevSpace = false
			excerpt += string(ch)
			continue
		}

		if !isPrevSpace {
			wordCount++
		}

		if wordCount == config.ExcerptSize {
			return excerpt + "..."
		}

		isPrevSpace = true
		excerpt += string(ch)
	}

	return excerpt
}

func (this Post) ExcerptToHtml() (string, error) {
	return markdown.ToString(this.Excerpt())
}
