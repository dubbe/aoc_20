package main

import (
	"flag"
	"fmt"
	"os/exec"
)

func main() {
	dayPtr := flag.Int("day", 0, "prepare for day")
	flag.Parse()

	fmt.Println("dayPtr:", *dayPtr)

	cmd := exec.Command("cp", "base", fmt.Sprintf("day%d", *dayPtr), "-r")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

