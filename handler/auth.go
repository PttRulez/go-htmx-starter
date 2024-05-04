package handler

import (
	"goth-anthonygg/view/auth"
	"net/http"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) error {
	return auth.Signin().Render(r.Context(), w)
}
