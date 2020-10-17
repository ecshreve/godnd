package genny

import (
	"bytes"
	"strings"
)

func snakeToCamel(s string) string {
	b := []byte(strings.TrimSpace(s))
	buffer := make([]byte, 0, len(s))

	tokens := bytes.Split(b, []byte("_"))
	for _, tok := range tokens {
		buffer = append(buffer, bytes.ToUpper(tok[:1])...)
		buffer = append(buffer, tok[1:]...)
	}

	return string(buffer)
}
