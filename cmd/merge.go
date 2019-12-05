package cmd

import (
	"flag"
	"fmt"
	"github.com/tk103331/jacocogo/core/tools"
	"os"
	"strings"
)

type mergeArgs struct {
	execFiles string
	destFile  string
}

type mergeCmd struct {
	args mergeArgs
}

func init() {
	add(&mergeCmd{})
}

func (mc *mergeCmd) name() string {
	return "merge"
}

func (mc *mergeCmd) desc() string {
	return "Merges multiple exec files into a new one."
}

func (mc *mergeCmd) parse(args []string) {
	mergeArgs := mergeArgs{}
	mergeFlagSet := flag.NewFlagSet(mc.name(), flag.ExitOnError)
	mergeFlagSet.StringVar(&mergeArgs.execFiles, "execfiles", "", "list of JaCoCo *.exec files to read")
	mergeFlagSet.StringVar(&mergeArgs.destFile, "destfile", "jacoco.exec", "file to write merged execution data to (default jacoco.exec)")
	err := mergeFlagSet.Parse(args)
	if err != nil {
		mergeFlagSet.PrintDefaults()
		return
	}
	mc.args = mergeArgs
}

func (mc *mergeCmd) exec() error {
	execfiles := mc.args.execFiles
	destfile := mc.args.destFile
	paths := strings.Split(execfiles, ",")

	loader := tools.NewFileLoader()

	for _, p := range paths {
		f, err := os.Open(p)
		if err != nil {
			fmt.Printf("open file error : %s", err.Error())
		}
		fmt.Printf("loading file %s ...", p)
		err = loader.Load(f)
		if err != nil {
			fmt.Printf("load file error : %s", err.Error())
		}
	}
	f, err := os.Create(destfile)
	if err != nil {
		fmt.Printf("create file error : %s", err.Error())
	}
	err = loader.Save(f)
	if err != nil {
		fmt.Printf("save file error : %s", err.Error())
	}
	return err
}
