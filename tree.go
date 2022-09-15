package egin

import (
	"bytes"

	"github.com/ericluj/egin/internal/bytesconv"
)

var (
	strColon = []byte(":")
	strStar  = []byte("*")
	strSlash = []byte("/")
)

type methodTree struct {
	method string
	root   *node
}

type methodTrees []methodTree

func (trees methodTrees) get(method string) *node {
	for _, tree := range trees {
		if tree.method == method {
			return tree.root
		}
	}
	return nil
}

type node struct {
	path     string
	fullPath string
}

func (n *node) addRoute(path string, handlers HandlersChain) {

}

type Param struct {
	Key   string
	Value string
}

type Params []Param

type skippedNode struct {
	path        string
	node        *node
	paramsCount int16
}

// 计算:和*个数
func countParams(path string) uint16 {
	var n uint16
	s := bytesconv.StringToBytes(path)
	n += uint16(bytes.Count(s, strColon))
	n += uint16(bytes.Count(s, strStar))
	return n
}

// 计算/个数
func countSections(path string) uint16 {
	s := bytesconv.StringToBytes(path)
	return uint16(bytes.Count(s, strSlash))
}
