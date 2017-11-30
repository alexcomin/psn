package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func pick(arg []string) string {
	var str string
	if arg[1] == "--all" {
		str = "ps aux | grep \"node \""
	} else if arg[1] == "--name" {
		if len(arg) < 3 {
			fmt.Println("missing parameter")
		} else {
			str = "ps aux | grep " + arg[2]
		}
	} else {
		str = "ps aux | grep \"node " + arg[1] + "\""
	}
	return str
}

func start() {
	arg := os.Args
	if len(arg) < 2 {
		fmt.Println("missing parameter")
		return
	}
	cmd := pick(arg)

	Cmd := exec.Command("bash", "-c", cmd)
	cmdOut, error := Cmd.Output()
	if error != nil {
		panic(error)
	}
	parse := strings.Split(string(cmdOut), "\n")
	for i := 0; i < len(parse)-1; i++ {
		math, _ := regexp.MatchString("bash|grep|psn", parse[i])
		if math != true {
			var arrayProc []string
			proc := strings.Split(parse[i], " ")
			for j := 0; j < len(proc); j++ {
				if proc[j] != "" {
					arrayProc = append(arrayProc, proc[j])
				}
			}
			arrayProc[4] = mb(arrayProc[4])
			arrayProc[5] = mb(arrayProc[5])
			pid := arrayProc[1]
			pwdx := exec.Command("pwdx", pid)
			pwdxOut, _ := pwdx.Output()
			fmt.Println(strings.Join(arrayProc, " "))
			color.Set(color.FgHiYellow)
			fmt.Println("Locate process", string(pwdxOut))
			color.Unset()
		}
	}
}

func mb(arg string) string {
	s, _ := strconv.ParseInt(arg, 10, 64)
	result := strconv.FormatInt(int64(s/1024), 10) + "Mib"
	return result
}

func main() {
	start()
}
