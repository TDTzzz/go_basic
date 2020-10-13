package gee

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			//首字符是通配符则跳出
			if item[0] == '*' {
				break
			}
		}
	}

	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	key := method + "-" + pattern

	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}

	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler

}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}
	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
	}

	return n, params
}

func (r *router) getRoutes(method string) []*node {
	root, ok := r.roots[method]
	if !ok {
		return nil
	}

	nodes := make([]*node, 0)
	root.travel(&nodes)
	return nodes
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)

	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)

	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}

//遍历children
func (r *router) WatchChildren() {

	for k, v := range r.handlers {
		log.Println(k, v)
	}
	for k, v := range r.roots {
		log.Println("遍历root", k, v)
		traversalChildren(v.children, 0)
	}
}

func traversalChildren(children []*node, height int) {
	if children == nil {
		log.Println("end")
	} else {
		for _, v := range children {
			fmt.Println("children--", "pattern:"+v.pattern, "---", "part:"+v.part, "---", height)

			fmt.Printf("%c[0;47;32m%s%s%c[0m\n", 0x1B, "pattern  ", v.pattern, 0x1B)
			fmt.Printf("%c[0;47;32m%s%s%c[0m\n", 0x1B, "part  ", v.part, 0x1B)
			fmt.Printf("%c[0;47;32m%s%d%c[0m\n", 0x1B, "height  ", height, 0x1B)
			fmt.Println()
			traversalChildren(v.children, height+1)
		}
	}
}
