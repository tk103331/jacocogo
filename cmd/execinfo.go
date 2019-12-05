package cmd

import (
	"flag"
	"fmt"
	"github.com/tk103331/jacocogo/core/data"
	"os"
	"strconv"
	"strings"
	"time"
)

type execInfoArgs struct {
	execFiles string
}

type execInfoCmd struct {
	args *execInfoArgs
}

type sessionVisitor struct {
}

func (sessionVisitor) VisitSessionInfo(info data.SessionInfo) error {
	fmt.Printf("Session \"%s\": %s - %s\n", info.Id, time.Unix(info.Start/1000, 0).String(), time.Unix(info.Dump/1000, 0).String())
	return nil
}

type executionVisitor struct {
}

func (executionVisitor) VisitExecutionData(data data.ExecutionData) error {
	count := 0
	for _, p := range data.Probes {
		if p {
			count++
		}
	}

	fmt.Printf("%16s  %3d of %3d   %s\n", hexStr(data.Id), count, len(data.Probes), data.Name)
	return nil
}

const pb = "1000000000000000000000000000000000000000000000000000000000000000"
const ph = "0000000000000000"

func hexStr(value int64) string {
	if value >= 0 {
		h := strconv.FormatInt(value, 16)
		return strings.Repeat("0", 16-len([]rune(h))) + h
	} else {
		b := []rune(strconv.FormatInt(9223372036854775807+value+1, 2))
		bin := append([]rune(pb)[0:64-len(b)], b...)
		ret := ""
		for i := 0; i < len(bin); i = i + 4 {
			tmp, _ := strconv.ParseInt(string(bin[i:i+4]), 2, 16)
			ret = ret + strconv.FormatInt(tmp, 16)
		}
		return ret
	}
}

func init() {
	add(&execInfoCmd{})
}

func (ec *execInfoCmd) name() string {
	return "execinfo"
}

func (ec *execInfoCmd) desc() string {
	return "Print exec file content in human readable format."
}

func (ec *execInfoCmd) parse(args []string) {
	execInfoArgs := &execInfoArgs{}
	flagSet := flag.NewFlagSet(ec.name(), flag.ExitOnError)
	flagSet.StringVar(&execInfoArgs.execFiles, "execfiles", "", "list of JaCoCo *.exec files to read")
	err := flagSet.Parse(args)
	if err != nil {
		flagSet.PrintDefaults()
	}
	ec.args = execInfoArgs
}

func (ec *execInfoCmd) exec() error {

	execfiles := ec.args.execFiles
	paths := strings.Split(execfiles, ",")

	for _, p := range paths {
		dump(p)
	}

	return nil
}

func dump(path string) {
	fmt.Printf("[INFO] Loading exec file %s.\n", path)
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("Loading exec file error: %s", err.Error())
	}
	defer f.Close()
	fmt.Println("CLASS ID         HITS/PROBES   CLASS NAME")
	reader := data.NewReader(f)
	reader.SetSessionVisitor(sessionVisitor{})
	reader.SetExecutionVisitor(executionVisitor{})
	reader.Read()
	fmt.Println()
}
