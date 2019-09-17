package jwtmiddleware

import (
	"net/http"
	"encoding/json"
	"INSTALLAMS/security"
	jwt "github.com/dgrijalva/jwt-go"
)

func JWTmiddleware(handler http.Handler) http.Handler{
	return http.HandlerFunc( func(response http.ResponseWriter, request *http.Request){
		stringToken := request.Header.Get("token")
		if stringToken == ""{
			respondWithError(response,http.StatusUnauthorized, "Unauthorized")
		}else{
			result, err:= jwt.Parse(stringToken, func(token *jwt.Token)(interface{}, error){
				return []byte(security.PrivateKey),nil
			})
			if err !=nil{
					respondWithError(response,http.StatusUnauthorized,err.Error())
			}else{
				if result.Valid{
					handler.ServeHTTP(response,request)
				}else{
					respondWithError(response,http.StatusUnauthorized, "InValid Key")
				}
			}
		}
	})
}

func respondWithError(w http.ResponseWriter, code int, msg string){
	respondWithJSON(w, code, map[string]string{"error":msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}){
	response,_:= json.Marshal(payload)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(code)
	w.Write(response)
}