package handlers

import (
	"INSTALLAMS/action"
	"INSTALLAMS/processfile"
	"INSTALLAMS/entities"
	"fmt"
)

//struct for AMS infomation

const (
	originPath        = `.\cmd\origin\`
	modifyPath           = `.\cmd\modify\`
	rstatfile         = `readstatus.txt`
	rstagfile         = `readstage.txt`
	instamsfile       = `installams.txt`
	uninstamsfile     = `uninstallams.txt`
	createWSPfile     = `createWSP.txt`
	delWSPfile        = `deleteWSP.txt`
	existfolderfile   = `existfolder.txt`
	existscriptsfile  = `existscripts.txt`
	existautofile     = `existauto.txt`
	existservstatfile = `existservstatus.txt`
	existAMSfile      = `existAMSServer.txt`
	defaultplugin     = `defaultplugin.txt`
	testpath          = originPath + `testssh.txt`
	sshtool           = `.\tools\putty\PLINK.exe `
	wspAMS            = `:/ams_build/AMSwsp/`
	buildpath         = `\\192.168.72.21\build`
	wspscripts        = `.\workspace\scripts`
	wspauto           = `.\workspace\auto`
	wspserverstatus   = `.\workspace\serverstatus`
	plugin			  = originPath + `plugin.txt`
	pscptool          = `.\tools\putty\PSCP.EXE -P 22 -q -pw `
	pscptoolfoler     = `.\tools\putty\PSCP.EXE -P 22 -r -q -pw `
)

func replaceStr(Path, replstr, fnstr string) string {
	cmd1 := `powershell -Command "(gc ` + Path + `) -replace '` + replstr + `', '` + fnstr + `' | Out-File -encoding ASCII ` + Path + `"`
	return cmd1
}
func replaceFile(modifyPath string, AMS entities.AMS) {
	cmd1 := replaceStr(modifyPath, "AMSVersion", AMS.AMSVersion) //change version
	action.ExecutePS(cmd1)
	cmd2 := replaceStr(modifyPath, "AMSRevision", AMS.AMSRevision) //change revision
	action.ExecutePS(cmd2)
}

func ModifyFile(AMS entities.AMS, originalPath, modifyPath string) {
	action.ExecuteCmd(`copy ` + originalPath + ` ` + modifyPath)
	replaceFile(modifyPath, AMS)
}
func ModifyFolder(AMS entities.AMS, originalFD, fixFD string) {
	var filecmds = map[int]string{
		1:  instamsfile,
		2:  uninstamsfile,
		3:  existfolderfile,
		4:  createWSPfile,
		5:  rstatfile,
		6:  rstagfile,
		7:  existscriptsfile,
		8:  existautofile,
		9:  existservstatfile,
		10: existAMSfile,
		11: defaultplugin,
	}
	for _, filecmd := range filecmds {
		ModifyFile(AMS, originalFD+filecmd, fixFD+filecmd)
	}
}

func CreateWSP(AMS entities.AMS) {
	ModifyFolder(AMS, originPath, modifyPath)
}

func DeleteWSP(AMS entities.AMS) {
	deleteWSP := modifyPath + delWSPfile
	deleteWSPcmd := SupportCmd(AMS, deleteWSP)
	action.ExecuteCmd(deleteWSPcmd)

}
func sshtoolfunc(AMS entities.AMS) string {
	return sshtool + `-ssh ` + AMS.AMSUsername + `@` + AMS.AMSServer + ` -P 22 -pw ` + AMS.AMSPassword + ` -m `
}

func AMSCmd(AMS entities.AMS, filename string) string {
	sshtoolcmd := sshtoolfunc(AMS)
	cmd := modifyPath + AMS.AMSServer + `\` + filename
	return sshtoolcmd + cmd
}

func SupportCmd(AMS entities.AMS, filename string) string {
	sshtoolcmd := sshtoolfunc(AMS)
	return sshtoolcmd + filename
}

func CreateWSPWindow(AMS entities.AMS) {
	fmt.Println("Enter CreateWSPWindow Func")
	windowWSP := modifyPath + AMS.AMSServer + `\`
	if exist := processfile.DirExists(windowWSP); !exist {
		processfile.MakeDir(windowWSP)
		ModifyFolder(AMS, originPath, windowWSP)
	}
	fmt.Println("Exit CreateWSPWindow Func")
}

func CreateWSPServer(AMS entities.AMS) {
	fmt.Println("Enter CreateWSPServer Func")
	existscriptfd := AMSCmd(AMS, existscriptsfile)
	if exist := action.ExecuteCmdOutput(existscriptfd); exist != "found\n" {
		des := AMS.AMSUsername + `@` + AMS.AMSServer + wspAMS
		SendFolder(wspscripts, des)
	}

	existautofd := AMSCmd(AMS, existautofile)
	if exist := action.ExecuteCmdOutput(existautofd); exist != "found\n" {
		des := AMS.AMSUsername + `@` + AMS.AMSServer + wspAMS
		SendFolder(wspauto, des)
	}
	existserverstatfd := AMSCmd(AMS, existservstatfile)
	if exist := action.ExecuteCmdOutput(existserverstatfd); exist != "found\n" {
		des := AMS.AMSUsername + `@` + AMS.AMSServer + wspAMS
		tool:=getPSCPFolder(AMS)
		SendFolder(tool+wspserverstatus, des)
	}
	fmt.Println("Exit CreateWSPServer Func")
}

func GetStatus(AMS entities.AMS) string {
	sshreadstatus := AMSCmd(AMS, rstatfile)
	return action.ExecuteCmdOutput(sshreadstatus)
}
func GetStage(AMS entities.AMS) string {
	sshreadstage := AMSCmd(AMS, rstagfile)
	return action.ExecuteCmdOutput(sshreadstage)
}

func GetLinkDownloadAMS(AMS entities.AMS) string {
	return `https://`+AMS.AMSServer+`:8443/ams-client/`
}


func SendFile(Source, Destination string) {
	sendfilecmd :=Source + ` ` + Destination
	action.ExecuteSSHCmd(sendfilecmd)
}

func SendFolder(Source, Destination string) {
	sendfoldercmd := Source + ` ` + Destination
	action.ExecuteSSHCmd(sendfoldercmd)
}

func getPSCP(AMS entities.AMS) string{
	return pscptool + AMS.AMSPassword +` `
}

func getPSCPFolder(AMS entities.AMS) string{
	return pscptoolfoler + AMS.AMSPassword +` `
}
