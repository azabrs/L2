package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"log"
	go_ps "github.com/mitchellh/go-ps"
)

func Cd(paths string){
	if err := os.Chdir(paths); err != nil{
		log.Println(err)
	}
	
}

func Pwd(){
	path, err := os.Getwd()
	if err != nil{
		log.Println(err)
	}
	fmt.Println(path)
}

func Echo(arg string) {
	arg = strings.Replace(arg, "echo ", "", 1)
	fmt.Println( arg)
}

func Ps(){
	proc, err := go_ps.Processes()
	if err != nil{
		log.Println(err)
	}
	fmt.Println("Pid PPid Executable")
	for _, p := range(proc){
		fmt.Println(p.Pid(), p.PPid(), p.Executable())
	}
}

func Kill(args []string){
	if len(args) < 2{
		log.Println("not enough arguments")
	}
	id, err := strconv.Atoi(args[1])
	if err != nil{
		log.Println(err)
	}
	proc, err := os.FindProcess(id)
	if err != nil{
		log.Println(err)
	}
	err = proc.Kill()
	if err != nil{
		log.Println(err)
	}
}

func Exec(args []string) {
	cmd := exec.Command(args[1], args[2:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil{
		log.Println(err)
	}
}
func ParseCom() error{
	reader := bufio.NewReader(os.Stdin)
	for{
		buf, err := reader.ReadString('\n')
		buf = strings.Trim(buf, "\n")
		if err != nil{
			return err
		}
		cmds := strings.Split(buf, " ")
		switch cmds[0]{
		case "quit":
			return nil
		case "cd":
			Cd(cmds[1])
		case "kill":
			Kill(cmds)
		case "ps":
			Ps()
		case "exec":
			Exec(cmds)
		case "pwd":
			Pwd()
		case "echo":
			Echo(buf)
		}
	}

}

func main(){
	err := ParseCom()
	if err != nil{
		log.Fatal(err)
	}
}