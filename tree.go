package egin

type methodTree struct {
}

type methodTrees []methodTree

type Param struct {
	Key   string
	Value string
}

type Params []Param

type node struct {
}

type skippedNode struct {
	path        string
	node        *node
	paramsCount int16
}
