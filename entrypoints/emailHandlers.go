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
	err:= domain.SendMail(ctx, email)

	if(err != nil){
		errorResponse(w,err)
		return
	}
	
	w.Write([]byte("Se ha mandado el correo"))
	

}
