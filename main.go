package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func execCommand(input string) error {
	input = strings.TrimSuffix(input, "\n")

	cmd := exec.Command(input)
	fmt.Println("cmd", cmd)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		// words := strings.Fields(input)
		handleErr(err)
		if err := execCommand(input); err != nil {
			fmt.Fprintln(os.Stdout, err)
		}
		fmt.Println(input)
		os.Exit(-1)
	}
}
