package handlers

import (
	"fmt"
	"net/http"
	"INSTALLAMS/entities"
	"encoding/json"
)

func StatusAms(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ENTERED StatusAMS")
	var amsserver entities.AMS
	err:= json.NewDecoder(r.Body).Decode(&amsserver)
	CreateWSPWindow(amsserver)
	if err!=nil{
		respondWithError(w,http.StatusBadRequest,err.Error())
	}else{
		status:=GetStatus(amsserver)
		respondWithJSON(w, http.StatusOK, map[string]string{"statusams":status})
	}	
	fmt.Println("EXITED StatusAMS")
}
	
