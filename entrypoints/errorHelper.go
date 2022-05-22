package entrypoints

import (
	"fmt"
	"net/http"
)

func errorResponse(w http.ResponseWriter, e error){
	fmt.Println(e)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte( e.Error() ))
}