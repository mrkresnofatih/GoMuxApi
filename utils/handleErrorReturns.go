package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func HandleErrorReturns(err error, w http.ResponseWriter) {
	log.Println("[Error] ", err.Error())
	w.WriteHeader(http.StatusBadRequest)
	f, _ := json.Marshal(map[string]string{
		"Message": err.Error(),
	})
	fmt.Fprintf(w, string(f))
}
