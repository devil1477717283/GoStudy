package gee

import "strings"

/*
参数匹配:。例如 /p/:lang/doc，可以匹配 /p/c/doc 和 /p/go/doc。
通配*。例如 /static/*filepath，可以匹配/static/fav.ico，也可以匹配/static/js/jQuery.js，这种模式常用于静态服务器，能够递归地匹配子路径。
*/
//node	一个节点就代表一个路由中的一部分
type node struct {
	pattern  string  //如果这是当前路由的最后一层则会被设置待匹配的路由，如果不是最后一层则会设置为空
	part     string  //路由中的一部分
	children []*node //子节点
	isWild   bool    //该部分是否能被精确匹配,part中含有*或者:时为true就可以进行模糊匹配，在注册路由的时候进行设置
}

//matchChild	第一个匹配成功的节点用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child //如果当前节点设置了可以被模糊匹配就可以直接返回当前节点，或者精确匹配成功之后返回
		}
	}
	return nil
}

//matchChildren	所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild { //精确匹配成功的或者模糊匹配成功的节点都返回
			nodes = append(nodes, child)
		}
	}
	return nodes
}

//insert	在树中插入节点
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern //匹配到最后一层才将pattern设置,其余层的pattern都为空
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'} //返回nil代表该路由还没注册，对路由进行注册，也就是对树进行插入
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") { //匹配到了最后一层,或者匹配到了*就终止
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
