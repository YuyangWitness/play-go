package router

import (
	"fmt"
	"net/http"
	"play-go/blue/context"
	"play-go/blue/trie"
	"play-go/blue/types"
	"strings"
)

type Router struct {
	root     map[string]*trie.Node
	handlers map[string]types.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		root:     make(map[string]*trie.Node),
		handlers: make(map[string]types.HandlerFunc),
	}
}

func (r *Router) AddRoute(method string, pattern string, handler types.HandlerFunc) {
	key := method + "-" + pattern
	_, ok := r.root[method]
	if !ok {
		r.root[method] = &trie.Node{}
	}
	parts := ParsePattern(pattern)
	r.root[method].Insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *Router) GetRoute(method string, pattern string) (*trie.Node, map[string]string) {
	searchParts := ParsePattern(pattern)
	params := make(map[string]string)
	root, ok := r.root[method]
	if !ok {
		return nil, nil
	}

	result := root.Search(searchParts, 0)

	if result != nil {
		parts := ParsePattern(result.Parttern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return result, params
	}

	return nil, nil
}

func (r *Router) Handler(c *context.Context) {
	n, params := r.GetRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.Parttern
		fmt.Println(r.handlers)
		fmt.Println(key)
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}

func ParsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := []string{}

	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}
