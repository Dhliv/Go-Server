package router

import (
	"net/http"
)

type Router struct {
	Rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		Rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) NewHandler(path, method string, handler http.HandlerFunc) {
	_, existPath := r.Rules[path]

	if !existPath {
		r.Rules[path] = make(map[string]http.HandlerFunc)
	}

	r.Rules[path][method] = handler
}

func (r *Router) FindHandler(path, method string) (http.HandlerFunc, bool, bool) {
	_, existsPath := r.Rules[path]
	if !existsPath { // si no existe la ruta, se notifica eso
		return nil, false, false
	}

	handler, existsMethod := r.Rules[path][method]
	return handler, existsPath, existsMethod
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, existsPath, existsMethod := r.FindHandler(request.URL.Path, request.Method)

	if !existsPath {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !existsMethod {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler(w, request)
}
