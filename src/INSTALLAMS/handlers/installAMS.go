package handlers

import (
	"INSTALLAMS/action"
	"fmt"
	"net/http"
	"time"
	"INSTALLAMS/entities"
	"encoding/json"
)

//InstallAMS function
func InstallAMS(w http.ResponseWriter, r *http.Request) {
	var amsserver entities.AMS
	err:= json.NewDecoder(r.Body).Decode(&amsserver)
	if err!=nil{
		respondWithError(w,http.StatusBadRequest,err.Error())
	}else{
		go installams(amsserver)
	}	
}
func installams(AMS entities.AMS) {
	CreateWSPWindow(AMS)
	CheckServer(AMS)
	Install(AMS)
}

//install
func Install(AMS entities.AMS) {
	fmt.Println("Enter Install Func")
	sshtestcmd := SupportCmd(AMS, testpath)
	installAMS := AMSCmd(AMS, instamsfile)
	tnow := time.Now()
	timediff := float64(61)
	for {
		if timediff > 60 {
			status:=GetStatus(AMS)
			errorssh := action.ExecuteSSHCmd(sshtestcmd)
			if errorssh != nil {
				fmt.Println("Connection not good! Please check the server AMS or use iptables -F")
				fmt.Println(errorssh)
			} else {
				CreateWSPServer(AMS)
				stage := GetStage(AMS)
				switch status {
				case "Disabled\n", "NotExist\n":
					fmt.Printf("AMS Server is %s", status)
					if stage != "BINRUNNING\n" && stage != "STARTING\n" && stage != "UNINSTALLING\n" && stage != "STOPPING\n" {
						action.ExecuteCmd(installAMS)
					}
				case "Running\n":
					if stage == "STARTED\n" {
						action.ExecuteCmd(installAMS)
					}
					if stage == "CREATED\n" {
						fmt.Println("Create First User Succesfully")
						action.Openbrowser(GetLinkDownloadAMS(AMS))
						return
					}
				case "WaitingStartup\n":
					if stage == "WaitingStartup\n" {
						//inform that lack of space, cannot install ams
					}
				}
			}
			tnow = time.Now()
		}
		timediff = time.Since(tnow).Seconds()
	}
	fmt.Println("Exit Install Func")
}


