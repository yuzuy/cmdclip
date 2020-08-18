package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/yuzuy/cmdclip"
)

func main() {
	os.Exit(run())
}

func run() int {
	flag.Parse()

	cmd := flag.Args()
	if len(cmd) == 0 {
		return report(errors.New("please specify a command"))
	}
	if err := cmdclip.Clip(cmd[0], cmd[1:]...); err != nil {
		return report(err)
	}

	return 0
}

func report(err error) int {
	fmt.Printf("cmdclip: %s\n", err)
	return 1
}
