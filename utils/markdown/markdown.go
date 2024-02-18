package markdown

import (
	"bytes"

	mark "github.com/yuin/goldmark"
)

func ToString(input string) (string, error) {
	var buffer bytes.Buffer

	if err := mark.Convert([]byte(input), &buffer); err != nil {
		return "", err
	}

	return buffer.String(), nil
}
