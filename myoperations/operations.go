package myoperations

import "mycontainers/node"

type Operations interface {
	Push(int) *node.Node
	Pop() (int, bool)
}
