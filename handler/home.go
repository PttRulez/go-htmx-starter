package handler

import (
	"goth-anthonygg/view/home"
	"net/http"
)

func HandlerHomeIndex(w http.ResponseWriter, r *http.Request) error {
	return home.Index().Render(r.Context(), w)
}
