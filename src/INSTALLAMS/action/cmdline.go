package action

import (
	"fmt"
	"log"
	"os/exec"
)

func ExecuteCmd(cmd string) {
	c := exec.Command("cmd", "/C", cmd)
	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}
func ExecuteCmdOutput(cmd string) string {
	out, err := exec.Command("cmd", "/C", cmd).Output()
	var outstr string = ""
	if err != nil {
		fmt.Println("Error: ", err)
		return outstr
	}
	outstr = string(out)
	return outstr
}
func ExecuteSSHCmd(cmd string) error {
	c := exec.Command("cmd", "/C", cmd)
	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
		return err
	}
	return nil
}
func ExecutePS(cmd string) {
	c := exec.Command("powershell", "-c", cmd)
	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}

func Openbrowser(url string) {
	var err error
	err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	if err != nil {
		log.Fatal(err)
	}
}
