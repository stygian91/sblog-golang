package models

import (
	"bytes"
	"sblog/config"
	"time"
	"unicode"

	mark "github.com/yuin/goldmark"
)

type Post struct {
	Id        uint
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (this Post) ContentToHtml() (string, error) {
	var buffer bytes.Buffer
	err := mark.Convert([]byte(this.Content), &buffer)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
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
			return excerpt
		}

		isPrevSpace = true
		excerpt += string(ch)
	}

	return excerpt
}
