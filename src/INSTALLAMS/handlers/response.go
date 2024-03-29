package handlers

import (
	"net/http"
	"encoding/json"
)

func respondWithError(w http.ResponseWriter, code int, msg string){
	respondWithJSON(w, code, map[string]string{"error":msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}){
	response,_:= json.Marshal(payload)
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.WriteHeader(code)
	w.Write(response)
}