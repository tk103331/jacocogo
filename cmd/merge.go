package cmd

import (
	"flag"
	"fmt"
	"github.com/tk103331/jacocogo/core/tools"
	"os"
	"strings"
)

type MergeArgs struct {
	execfiles string
	destfile  string
}

type MergeCmd struct {
	args MergeArgs
}

func NewMergeCmd(args []string) *MergeCmd {
	return (&MergeCmd{}).parse(args)
}

func (mc *MergeCmd) parse(args []string) *MergeCmd {
	mergeArgs := MergeArgs{}
	mergeFlagSet := flag.NewFlagSet("merge", flag.ExitOnError)
	mergeFlagSet.StringVar(&mergeArgs.execfiles, "execfiles", "", "list of JaCoCo *.exec files to read")
	mergeFlagSet.StringVar(&mergeArgs.destfile, "destfile", "jacoco.exec", "file to write merged execution data to (default jacoco.exec)")
	err := mergeFlagSet.Parse(args)
	if err != nil {
		mergeFlagSet.PrintDefaults()
		return nil
	}
	mc.args = mergeArgs
	return mc
}

func (mc *MergeCmd) Exec() {
	execfiles := mc.args.execfiles
	destfile := mc.args.destfile
	paths := strings.Split(execfiles, ",")

	loader := tools.NewFileLoader()

	for _, p := range paths {
		f, err := os.Open(p)
		if err != nil {
			fmt.Printf("load file error : %s", err.Error())
		}
		loader.Load(f)
	}
	f, err := os.Create(destfile)
	if err != nil {
		fmt.Printf("save file error : %s", err.Error())
	}
	loader.Save(f)
}
