package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

var ErrNoPath = errors.New("no path provided")

func execCommand(input string) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")
	// fmt.Println(args)

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			writeToLog("error.log", ErrNoPath.Error())
			return ErrNoPath
		}
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}
	cmd := exec.Command(args[0], args[1:]...)
	// fmt.Println("cmd", cmd)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	user, _ := user.Current()

	for {
		dir, err := os.Getwd()
		if err != nil {
			writeToLog("error.log", err.Error())
			handleErr(err)
		}
		fmt.Printf("%s:%s:> ", user.Username, dir)
		input, err := reader.ReadString('\n')
		// words := strings.Fields(input)
		if err != nil {
			writeToLog("error.log", err.Error())
			handleErr(err)
		}
		if err := execCommand(input); err != nil {
			writeToLog("error.log", err.Error())
			fmt.Fprintln(os.Stdout, err)
		}
	}
}

func writeToLog(fname string, data string) {
	f, err := os.OpenFile(fname, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		handleErr(err)
	}
	defer f.Close()
	data = fmt.Sprintf("%s \n", data)
	_, err = f.WriteString(data)
}
