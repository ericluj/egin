package egin

import (
	"math"
	"net/http"
	"net/url"
)

// TODO:
const abortIndex int8 = math.MaxInt8 >> 1

type Context struct {
	engine *Engine

	writermem responseWriter
	Request   *http.Request

	Writer       ResponseWriter
	Params       Params
	handlers     HandlersChain
	index        int8
	fullPath     string
	Keys         map[string]any
	Errors       errorMsgs
	Accepted     []string
	queryCache   url.Values
	formCache    url.Values
	sameSite     http.SameSite
	params       *Params // TODO:
	skippedNodes *[]skippedNode
}

type ResponseWriter interface {
}

func (c *Context) reset() {
	c.Writer = &c.writermem
	c.Params = c.Params[:0]
	c.handlers = nil
	c.index = -1
	c.fullPath = ""
	c.Keys = nil
	c.Errors = c.Errors[:0]
	c.Accepted = nil
	c.queryCache = nil
	c.formCache = nil
	c.sameSite = 0
	*c.params = (*c.params)[:0]
	*c.skippedNodes = (*c.skippedNodes)[:0]
}

func (c *Context) Next() {

}
