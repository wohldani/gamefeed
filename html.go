package gamefeed

import (
	"golang.org/x/net/html"
	"io"
)

func findNodes(content io.Reader, node string) []html.Token {
	tokens := html.NewTokenizer(content)
	var resp []html.Token
	for {
		tokenType := tokens.Next()
		switch {
		case tokenType == html.ErrorToken:
			return resp
		case tokenType == html.StartTagToken:
			token := tokens.Token()
			if token.Data == node {
				resp = append(resp, token)
			}
		}
	}
	return resp
}
