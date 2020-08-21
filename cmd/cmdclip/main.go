package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/yuzuy/cmdclip"
)

var (
	pFlag = flag.Bool("p", false, "Print output of the command")
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
	out, err := cmdclip.Clip(cmd[0], cmd[1:]...)
	if err != nil {
		return report(err)
	}
	if *pFlag {
		fmt.Println(out)
	}

	return 0
}

func report(err error) int {
	fmt.Printf("cmdclip: %s\n", err)
	return 1
}
