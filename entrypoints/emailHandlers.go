package entrypoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	"personal-page-back/domain"
)

func PostEmail(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Post email handler")
	ctx := req.Context()
	setHeaders(w)

	var email domain.Email
	json.NewDecoder(req.Body).Decode(&email)
	domain.SendMail(ctx, email)
	w.Write([]byte("Hola"))
}
