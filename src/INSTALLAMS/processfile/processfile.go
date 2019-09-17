package processfile

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func ReadFile(Filename string) []string {
	var lines []string
	file, err := os.Open(Filename)
	if err != nil {
		fmt.Println(err)
		//os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
func CreateFile(Filename string) {
	var file, err = os.Create(Filename)
	if isError(err) {
		return
	}
	defer file.Close()
}
func DeleteFile(Filename string) {
	var err = os.Remove(Filename)
	if isError(err) {
		return
	}
}
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func FindStr(Filename, Search string) (string, error) {
	var tempstr string = ""
	f, err := os.Open(Filename)
	if err != nil {
		fmt.Println("exit ", err)
		return tempstr, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	line := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), Search) {
			fmt.Println("find")
			temp := fmt.Sprintf("%d %s \t%s\n", line, Filename, scanner.Text())
			tempstr += temp
		}
		line++
	}

	if err := scanner.Err(); err != nil {
		return tempstr, err
	}
	return tempstr, nil
}
func WriteFile(Filename, Content string) {
	f, err := os.OpenFile(Filename, os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(Content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Write ", l, " bytes")
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
func GoFindStr(wg *sync.WaitGroup, m *sync.Mutex, Filename, Search, Fileresult string) {
	result, err := FindStr(Filename, Search)
	if err != nil {
		return
	}
	m.Lock()
	WriteFile(Fileresult, result)
	m.Unlock()
	wg.Done()
}
func Exists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
func DirExists(Directory string) bool {
	_, err := os.Stat(Directory)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func MakeDir(Directory string) bool {
	err := os.MkdirAll(Directory, 0644)
	if err != nil {
		return false
	}
	return true
}
