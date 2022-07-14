package main

type Tree interface{

}
type node struct{
	data int
	left *node
	right *node
}
type binaryTree struct {
	root *node
}

type multiWayTreeNode struct {
	data int
	child []*node
}
type multiWayTree struct {
	root *multiWayTreeNode
}