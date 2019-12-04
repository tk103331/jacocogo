package cmd

import (
	"flag"
	"fmt"
	"github.com/tk103331/jacocogo/core/tools"
	"os"
	"strconv"
)

type DumpArgs struct {
	Address  string
	Port     uint64
	DestFile string
	Reset    bool
	Retry    uint64
}

type DumpCmd struct {
	args    DumpArgs
	argsErr error
}

func NewDumpCmd(args []string) *DumpCmd {
	return (&DumpCmd{}).parse(args)
}

func (dc *DumpCmd) parse(args []string) *DumpCmd {
	dumpArgs := DumpArgs{}
	dumpFlagSet := flag.NewFlagSet("dump", flag.ExitOnError)
	dumpFlagSet.StringVar(&dumpArgs.Address, "address", "localhost", "host name or ip address to connect to (default localhost)")
	dumpFlagSet.Uint64Var(&dumpArgs.Port, "port", 6300, "the port to connect to (default 6300)")
	dumpFlagSet.StringVar(&dumpArgs.DestFile, "destfile", "jacoco.exec", "file to write execution data to (default jacoco.exec)")
	dumpFlagSet.BoolVar(&dumpArgs.Reset, "reset", false, "reset execution data on test target after dump (default false)")
	dumpFlagSet.Uint64Var(&dumpArgs.Retry, "retry", 10, "number of retries (default 10)")
	err := dumpFlagSet.Parse(args)
	if err != nil {
		dc.argsErr = err
		dumpFlagSet.PrintDefaults()
	}
	dc.args = dumpArgs
	return dc
}
func (dc *DumpCmd) Exec() error {
	if dc.argsErr != nil {
		return dc.argsErr
	}
	dumpArgs := dc.args
	client := tools.NewDumpClient()
	defer client.Close()
	client.SetDump(true)
	client.SetReset(dumpArgs.Reset)
	client.SetRetryCount(int(dumpArgs.Retry))
	client.OnConnecting = func(address string) {
		fmt.Printf("connecting to %s\n ...", address)
	}
	client.OnConnectFailed = func(err error) {
		fmt.Printf("connect failed : %s\n ...", err.Error())
	}
	loader, err := client.Dump(dumpArgs.Address + ":" + strconv.Itoa(int(dumpArgs.Port)))
	if err != nil {
		return err
	}
	file, err := os.Create(dumpArgs.DestFile)
	defer file.Close()
	if err != nil {
		return err
	}
	err = loader.Save(file)
	if err != nil {
		return err
	}
	return nil
}
