package service

import (
	"context"
	"fmt"
	"strings"
)

type HandlerGeter interface {
	HandleGET(ctx context.Context, word string, handler func(data string, path []string)) error
}

type cambridgeWordResolver struct {
	parser      HandlerGeter
	handlerFunc func(w *Word) func(string, []string)
}

type Word struct {
	Origin         string   `json:"origin"`
	Opposite       string   `json:"opposite"`
	Transcript     string   `json:"transcript"`
	Type           string   `json:"type"`
	Translate      []string `json:"translate"`
	Means          []string `json:"means"`
	TranslateMeans []string `json:"translate_means"`
	Level          string   `json:"level"`
}

func NewCambridgeWordResolver() *cambridgeWordResolver {
	parser := NewHandler("dictionary.cambridge.org", "dictionary/english-russian")
	handler := func(w *Word) func(string, []string) {
		return func(data string, path []string) {
			addr := strings.Join(path, "/")
			if !strings.HasPrefix(addr, "html/body/div/div/div/div/article/") {
				return
			}
			_, addr, _ = strings.Cut(addr, "html/body/div/div/div/div/article/")
			switch addr {
			case "div/div/div/div/div/div/div/span/span":
				w.Origin = strings.TrimSpace(fmt.Sprintf("%s %s", w.Origin, data))
			case "div/div/div/div/div/div/span/span/span":
				w.Transcript = strings.TrimSpace(fmt.Sprintf("%s %s", w.Transcript, data))
			case "div/div/div/div/div/div/div/span":
				w.Type = strings.TrimSpace(fmt.Sprintf("%s %s", w.Type, data))
			case "div/div/div/div/div/div/div/div/div/div/span":
				w.Translate = append(w.Translate, data)
			case "div/div/div/div/div/div/div/div/div/div/span/b":
				w.Means = append(w.Means, data)
			case "div/div/div/div/div/div/div/div/div/div/div/div/span":
				w.TranslateMeans = append(w.TranslateMeans, data)
			case "div/div/div/div/div/div/div/div/div/div/div/div/div/div/div/a/span",
				"div/div/div/div/div/div/div/div/div/div/div/div/div/div/div/a/span/span":
				w.Opposite = strings.TrimSpace(fmt.Sprintf("%s %s", w.Opposite, data))
			case "div/div/div/div/div/div/div/div/div/div/span/span":
				w.Level = strings.TrimSpace(fmt.Sprintf("%s %s", w.Level, data))
			}
		}
	}
	return &cambridgeWordResolver{
		parser:      parser,
		handlerFunc: handler,
	}
}

func (c *cambridgeWordResolver) GetWord(ctx context.Context, world string) (Word, error) {
	var w Word
	err := c.parser.HandleGET(ctx, world, c.handlerFunc(&w))
	if err != nil {
		return w, fmt.Errorf("error wile try get word %w", err)
	}
	return w, nil
}
