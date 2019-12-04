package tools

import (
	"errors"
	"github.com/tk103331/jacocogo/core/runtime"
	"net"
	"time"
)

type DumpClient struct {
	conn            net.Conn
	dump            bool
	reset           bool
	retryCount      int
	retryDelay      time.Duration
	OnConnecting    func(address string)
	OnConnectFailed func(err error)
}

func NewDumpClient() *DumpClient {
	return &DumpClient{nil, true, false, 0, 1000, nil, nil}
}

func (dc *DumpClient) SetDump(dump bool) {
	dc.dump = dump
}
func (dc *DumpClient) SetReset(reset bool) {
	dc.reset = reset
}
func (dc *DumpClient) SetRetryCount(retryCount int) {
	dc.retryCount = retryCount
}
func (dc *DumpClient) SetRetryDelay(retryDelay time.Duration) {
	dc.retryDelay = retryDelay
}
func (dc *DumpClient) Close() error {
	if dc.conn != nil {
		return dc.conn.Close()
	}
	return nil
}
func (dc *DumpClient) Dump(address string) (*FileLoader, error) {
	fileLoader := NewFileLoader()
	conn, err := dc.tryConnect(address)
	if err != nil {
		return fileLoader, nil
	}
	defer conn.Close()
	reader := runtime.NewControlReader(conn)
	writer := runtime.NewControlWriter(conn)

	reader.SetSessionVisitor(fileLoader.SessionStore())
	reader.SetExecutionVisitor(fileLoader.ExecutionStore())

	err = writer.VisitDumpCommand(dc.dump, dc.reset)
	if err != nil {
		return fileLoader, err
	}
	read, err := reader.Read()
	if err != nil {
		return fileLoader, err
	}
	if !read {
		return fileLoader, errors.New("socket closed unexpectedly")
	}
	return fileLoader, nil
}
func (dc *DumpClient) tryConnect(address string) (net.Conn, error) {
	count := 0
	for {
		dc.OnConnecting(address)
		conn, err := net.Dial("tcp", address)
		if err == nil {
			return conn, nil
		} else if count+1 > dc.retryCount {
			return nil, err
		}
		count = count + 1
		dc.OnConnectFailed(err)
		time.Sleep(dc.retryDelay)
	}
}
