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

type nodeType uint8

const (
	root nodeType = iota + 1
	param
	catchAll
)

type node struct {
	// 结点路径
	path string
	// 分支首字母拼接而成
	indices string
	// 结点是否是参数结点
	wildChild bool
	// 结点类型
	nType nodeType
	// 优先级（子结点等注册的handler数量）
	priority uint32
	// 子结点
	children []*node
	// 处理函数链
	handlers HandlersChain
	// 完整路径
	fullPath string
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func longestCommonPrefix(a, b string) int {
	i := 0
	max := min(len(a), len(b))
	for i < max && a[i] == b[i] {
		i++
	}
	return i
}

func (n *node) addRoute(path string, handlers HandlersChain) {
	fullPath := path
	n.priority++

	// 空树
	if len(n.path) == 0 && len(n.children) == 0 {
		n.insertChild(path, fullPath, handlers)
		n.nType = root
		return
	}

	parentFullPathIndex := 0

walk:
	for {
		// 获取公共前缀
		i := longestCommonPrefix(path, n.path)

		// 当前结点的一部分为公共前缀，将其后的部分分裂出来作为新的子结点
		if i < len(n.path) {
			child := node{
				path:      n.path[i:],
				wildChild: n.wildChild,
				indices:   n.indices,
				children:  n.children,
				handlers:  n.handlers,
				priority:  n.priority - 1,
				fullPath:  n.fullPath,
			}
			n.children = []*node{&child}
			n.indices = bytesconv.BytesToString([]byte{n.path[i]})
			n.path = path[:i]
			n.handlers = nil
			n.wildChild = false
			n.fullPath = fullPath[:parentFullPathIndex+i]
		}

		// 加入path的非公共前缀部分作为新的子结点
		if i < len(path) {
			path = path[i:]
			c := path[0]

			// TODO:
			if n.nType == param && c == '/' && len(n.children) == 1 {
				parentFullPathIndex += len(n.path)
				n = n.children[0]
				n.priority++
				continue walk
			}

		}
	}
}

func (n *node) insertChild(path string, fullPath string, handlers HandlersChain) {
	// TODO:

	n.path = path
	n.fullPath = fullPath
	n.handlers = handlers
}

type nodeValue struct {
	handlers HandlersChain
	params   *Params
	tsr      bool
	fullPath string
}

func (n *node) getValue(path string, params *Params, skippedNodes *[]skippedNode, unescape bool) (value nodeValue) {
	return
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
