package main

import (
	"flag"
	"github.com/tk103331/jacocogo/cmd"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		cmd.Usage()
		return
	}
	cmd.Exec(args[0], args[1:])
}
