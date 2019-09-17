package handlers

import (
	"INSTALLAMS/processfile"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"unicode"
	"INSTALLAMS/entities"
	"regexp"
	"INSTALLAMS/action"
)


func SelectAMS(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entered SelectAMS function")
	var rcmAMS = entities.RecommendAMS{
		Max:        0,
		AMSVersion: "0",
		AMSRevison:  "0",
	}
	NewAMSVersion := FindNewAMS(buildpath)
	rcmAMS=FindNewRevision(buildpath+NewAMSVersion+`\All_Build\`, NewAMSVersion)
	respondWithJSON(w, http.StatusOK, map[string]string{"version":rcmAMS.AMSVersion,"revision":rcmAMS.AMSRevison})
	fmt.Println("Exited SelectAMS function")
}



func CheckRevision(link string) []string {
	folder, err := os.Open(link)
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer folder.Close()
	lists, _ := folder.Readdirnames(0) // 0 to read all files and folders
	return lists
}

func FindNewAMS(link string) string {
	Versions, err := ioutil.ReadDir(link)
	if err != nil {
		log.Fatal(err)
	}
	NewAMSServer := ""
	for _, Version := range Versions {
		a := len(Version.Name())
		if isInt(Version.Name()[0]) && isInt(Version.Name()[a-1]) {
			NewAMSServer = Version.Name()
		}
	}
	return NewAMSServer
}


func CheckServer(AMS entities.AMS) {
	fmt.Println("Enter Checkserver Func")
	pattern := regexp.MustCompile(`\n|ams-|-`)
	existAMSServercmd := AMSCmd(AMS, originPath+existAMSfile)
	servers := action.ExecuteCmdOutput(existAMSServercmd)
	fmt.Println(servers)
	folders := pattern.Split(servers, -1)
	if len(folders) > 0 {
		for i := 1; i < len(folders); i += 3 {
			AMStemp := AMS
			AMStemp.AMSVersion = folders[i]
			AMStemp.AMSRevision = folders[i+1]
			if AMStemp.AMSVersion != AMS.AMSVersion || AMStemp.AMSRevision != AMS.AMSRevision {
				Uninstall(AMS)
			}
		}
	}
	fmt.Println(folders)
	fmt.Println("Exit Checkserver Func")
}

func  FindNewRevision(link, NewAMSVersion string) (rcmAMS entities.RecommendAMS) {
	Revisions, err := ioutil.ReadDir(link)
	if err != nil {
		log.Fatal(err)
	}
	count := 0
	for _, Revision := range Revisions {
		a := len(Revision.Name())
		if isInt(Revision.Name()[0]) && isInt(Revision.Name()[a-1]) {
			lists := CheckRevision(link + Revision.Name())
			files := processfile.ReadFile(plugin)

			for _, filename := range lists {
				for inc := 0; inc < len(files); inc++ {
					if strings.Contains(filename, files[inc]) {
						count++
						break
					}
				}
			}
		}
		fmt.Println("count:", count)
		if count >= rcmAMS.Max {
			rcmAMS.Max = count
			rcmAMS.AMSVersion = NewAMSVersion
			rcmAMS.AMSRevison = Revision.Name()
		}
		count = 0
	}
	return
}

func isInt(s byte) bool {
	if !unicode.IsDigit(rune(s)) {
		return false
	}
	return true
}