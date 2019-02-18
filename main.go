package main

import (
	"fmt"
	"os"

	"github.com/fbsb/dotf/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
