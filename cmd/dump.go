package cmd

import (
	"flag"
	"fmt"
	"github.com/tk103331/jacocogo/core/tools"
	"os"
	"strconv"
)

type dumpArgs struct {
	Address  string
	Port     uint64
	DestFile string
	Reset    bool
	Retry    uint64
}

type dumpCmd struct {
	args dumpArgs
}

func init() {
	cmd := &dumpCmd{}
	cmds[cmd.name()] = cmd
}

func (dc *dumpCmd) name() string {
	return "dump"
}

func (dc *dumpCmd) desc() string {
	return "Request execution data from a JaCoCo agent running in 'tcpserver' output mode."
}

func (dc *dumpCmd) parse(args []string) {
	dumpArgs := dumpArgs{}
	dumpFlagSet := flag.NewFlagSet("dump", flag.ExitOnError)
	dumpFlagSet.StringVar(&dumpArgs.Address, "address", "localhost", "host name or ip address to connect to (default localhost)")
	dumpFlagSet.Uint64Var(&dumpArgs.Port, "port", 6300, "the port to connect to (default 6300)")
	dumpFlagSet.StringVar(&dumpArgs.DestFile, "destfile", "jacoco.exec", "file to write execution data to (default jacoco.exec)")
	dumpFlagSet.BoolVar(&dumpArgs.Reset, "reset", false, "reset execution data on test target after dump (default false)")
	dumpFlagSet.Uint64Var(&dumpArgs.Retry, "retry", 10, "number of retries (default 10)")
	err := dumpFlagSet.Parse(args)
	if err != nil {
		dumpFlagSet.PrintDefaults()
	}
	dc.args = dumpArgs
}
func (dc *dumpCmd) exec() error {
	dumpArgs := dc.args
	client := tools.NewDumpClient()
	defer client.Close()
	client.SetDump(true)
	client.SetReset(dumpArgs.Reset)
	client.SetRetryCount(int(dumpArgs.Retry))
	client.OnConnecting = func(address string) {
		fmt.Printf("connecting to %s ...\n", address)
	}
	client.OnConnectFailed = func(err error) {
		fmt.Printf("connect failed : %s\n", err.Error())
	}
	loader, err := client.Dump(dumpArgs.Address + ":" + strconv.Itoa(int(dumpArgs.Port)))
	if err != nil {
		return err
	}
	file, err := os.Create(dumpArgs.DestFile)
	if err != nil {
		return err
	}
	defer file.Close()
	err = loader.Save(file)
	if err != nil {
		return err
	}
	return nil
}
