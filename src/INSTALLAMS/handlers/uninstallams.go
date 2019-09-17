package handlers

import (
	"INSTALLAMS/action"
	"fmt"
	"net/http"
	"time"
	"INSTALLAMS/entities"
	"encoding/json"
)

func UninstallAMS(w http.ResponseWriter, r *http.Request) {
	var amsserver entities.AMS
	err:= json.NewDecoder(r.Body).Decode(&amsserver)
	if err!=nil{
		respondWithError(w,http.StatusBadRequest,err.Error())
	}else{
		go uninstallams(amsserver)
	}	
}

func uninstallams(AMS entities.AMS) {
	CreateWSPWindow(AMS)
	Uninstall(AMS)
}
func Uninstall(AMS entities.AMS) {
	fmt.Println("Enter Uninstall Func")
	sshtestcmd := SupportCmd(AMS, testpath)
	uninstall := AMSCmd(AMS, uninstamsfile)
	fmt.Println(sshtestcmd)
	tnow := time.Now()
	timediff := float64(61)
	for {
		if timediff > 60 {
			errorssh := action.ExecuteSSHCmd(sshtestcmd)
			if errorssh != nil {
				fmt.Println("Connection not good! Please check the server AMS or use iptables -F")
				fmt.Println(errorssh)
			} else {
				//fmt.Println(workspace)
				CreateWSPServer(AMS)
				status:=GetStatus(AMS)
				stage := GetStage(AMS)
				switch status {
				case "NotExist\n":
					fmt.Printf("AMS Server's been uninstalled completely")
					return
					//DeleteWSP(AMS)
				case "Disabled\n":
					if stage != "UNINSTALLING\n" {
						action.ExecuteCmd(uninstall)
					}

				case "Running\n":
					if stage != "STOPPING\n" {
						action.ExecuteCmd(uninstall)
					}
				}
			}
			tnow = time.Now()
		}
		timediff = time.Since(tnow).Seconds()
	}
	fmt.Println("Exit Uninstall Func")

}
