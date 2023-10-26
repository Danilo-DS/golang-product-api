package routers

import "net/http"

type RouterModel struct {
	URI          string
	HasParameter bool
	Parameters   []string
	Method       string
	HandleFunc   func(http.ResponseWriter, *http.Request)
}
