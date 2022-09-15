package egin

import "net/http"

const (
	noWritten     = -1
	defaultStatus = http.StatusOK
)

type responseWriter struct {
	http.ResponseWriter
	size   int
	status int
}

func (w *responseWriter) reset(writer http.ResponseWriter) {
	w.ResponseWriter = writer
	w.size = noWritten
	w.status = defaultStatus
}
