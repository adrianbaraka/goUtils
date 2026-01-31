package main

import (
	"fmt"
	"strings"

	"github.com/adrianbaraka/goutils/cli"
)

func main() {
	cli.RunCmd(true, false, true, "ls", "-la")

	strings.Contains("Hello", "h")

	var hi bool

	if hi {
		fmt.Println("Hi")
	}


}
