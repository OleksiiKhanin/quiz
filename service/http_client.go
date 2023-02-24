package service

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
	"golang.org/x/net/html"
)

type Doer interface {
	Do(*retryablehttp.Request) (*http.Response, error)
}

type parser struct {
	u      url.URL
	client Doer
}

func NewHandler(addr, dictionary string) *parser {
	return &parser{
		u: url.URL{
			Scheme: "https",
			Host:   addr,
			Path:   dictionary,
		},
		client: retryablehttp.NewClient(),
	}
}

func (o *parser) HandleGET(ctx context.Context, word string, handler func(data string, path []string)) error {
	reqURL := o.u.JoinPath(word)
	req, err := retryablehttp.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/109.0")
	if err != nil {
		return err
	}
	resp, err := o.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	node, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}
	var f func(path []string, n *html.Node)
	f = func(path []string, n *html.Node) {
		for item := n.FirstChild; item != nil; item = item.NextSibling {
			s := strings.TrimSpace(item.Data)
			if item.Type == html.TextNode {
				if len(s) > 0 {
					handler(s, path)
				}
			} else {
				cpPath := make([]string, len(path))
				copy(cpPath, path)
				cpPath = append(cpPath, s)
				f(cpPath, item)
			}
		}
	}
	f([]string{}, node)
	return nil
}
