package service

import (
	"context"
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

type HandlerGeter interface {
	HandleGET(ctx context.Context, word string, handler func(data string, path []string, attr ...html.Attribute)) error
}

type cambridgeWorldResolver struct {
	parser      HandlerGeter
	handlerFunc func(data string, path []string, attr ...html.Attribute)
}

func NewCambridgeWorldResolver() *cambridgeWorldResolver {
	parser := NewHandler("dictionary.cambridge.org", "dictionary/english-russian")
	handler := func(data string, path []string, attr ...html.Attribute) {
		str := strings.TrimSpace(data)
		if len(str) > 0 {
			fmt.Fprintf(os.Stderr, "Path:%s,Data:%s\n", strings.Join(path, "/"), str)
		}
	}
	return &cambridgeWorldResolver{
		parser:      parser,
		handlerFunc: handler,
	}
}

func (c *cambridgeWorldResolver) GetWorld(ctx context.Context, world string) map[string]string {
	err := c.parser.HandleGET(ctx, world, c.handlerFunc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error %s wile try get word %s\n", err.Error(), world)
	}
	return nil
}
