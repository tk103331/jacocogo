package cmd

import (
	"flag"
	"fmt"
	"github.com/tk103331/jacocogo/core/tools"
	"os"
)

type dumpArgs struct {
	address  string
	port     uint64
	destFile string
	reset    bool
	retry    uint64
}

type dumpCmd struct {
	args dumpArgs
}

func init() {
	add(&dumpCmd{})
}

func (dc *dumpCmd) name() string {
	return "dump"
}

func (dc *dumpCmd) desc() string {
	return "Request execution data from a JaCoCo agent running in 'tcpserver' output mode."
}

func (dc *dumpCmd) parse(args []string) {
	dumpArgs := dumpArgs{}
	dumpFlagSet := flag.NewFlagSet(dc.name(), flag.ExitOnError)
	dumpFlagSet.StringVar(&dumpArgs.address, "address", "localhost", "host name or ip address to connect to (default localhost)")
	dumpFlagSet.Uint64Var(&dumpArgs.port, "port", 6300, "the port to connect to (default 6300)")
	dumpFlagSet.StringVar(&dumpArgs.destFile, "destfile", "jacoco.exec", "file to write execution data to (default jacoco.exec)")
	dumpFlagSet.BoolVar(&dumpArgs.reset, "reset", false, "reset execution data on test target after dump (default false)")
	dumpFlagSet.Uint64Var(&dumpArgs.retry, "retry", 10, "number of retries (default 10)")
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
	client.SetReset(dumpArgs.reset)
	client.SetRetryCount(int(dumpArgs.retry))
	client.OnConnecting = func(address string) {
		fmt.Printf("connecting to %s ...\n", address)
	}
	client.OnConnectFailed = func(err error) {
		fmt.Printf("connect failed : %s\n", err.Error())
	}
	loader, err := client.Dump(fmt.Sprintf("%s:%d", dumpArgs.address, dumpArgs.port))
	if err != nil {
		return err
	}
	file, err := os.Create(dumpArgs.destFile)
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
