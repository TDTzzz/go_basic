package gee

import "fmt"

type node struct {
	pattern  string  //待匹配路由
	part     string  //路由中的一部分
	children []*node //子节点
	isWild   bool    //是否是精确匹配
}

func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s,part%s,isWild=%t}", n.pattern, n.part, n.isWild)
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

}
