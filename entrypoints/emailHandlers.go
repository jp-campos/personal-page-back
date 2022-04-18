package entrypoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	"personal-page-back/domain"
)

func PostEmail(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	var email domain.Email
	json.NewDecoder(req.Body).Decode(&email)
	fmt.Println(email)
	domain.SendMail(ctx, email)
}
