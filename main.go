package main

import (
	"fmt"
	"github.com/tk103331/jacocogo/cmd"
	"os"
)

func main() {
	cmdName := os.Args[1]
	args := os.Args[2:]
	switch cmdName {
	case "dump":
		cmd.NewDumpCmd(args).Exec()
	case "merge":
		cmd.NewMergeCmd(args).Exec()
	default:
		fmt.Println("unsupported command")
		fmt.Println("supported commands: \n\tdump \n\tmerge ")
	}
}
