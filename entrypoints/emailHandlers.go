package entrypoints

import (
	"encoding/json"
	"net/http"
	"personal-page-back/domain"
)

func PostEmail(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	setHeaders(w)

	var email domain.Email
	json.NewDecoder(req.Body).Decode(&email)
	domain.SendMail(ctx, email)
	w.Write([]byte("Hola"))
}
