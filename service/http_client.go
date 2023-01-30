package service

import (
	"context"
	retryablehttp "github.com/hashicorp/go-retryablehttp"
	"golang.org/x/net/html"
	"net/http"
	"net/url"
	"strings"
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

func (o *parser) HandleGET(ctx context.Context, word string, handler func(data string, path []string, attr ...html.Attribute)) error {
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
			if item.Type == html.TextNode {
				handler(item.Data, path, item.Attr...)
			} else {
				if s := strings.TrimSpace(item.Data); len(s) > 0 {
					path = append(path, s)
				}
				f(path, item)
			}
		}
	}
	f([]string{}, node)
	return nil
}
