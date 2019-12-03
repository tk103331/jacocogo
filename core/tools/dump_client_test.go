package tools

import (
	"os"
	"testing"
)

func TestDumpClient_Dump(t *testing.T) {
	client := NewDumpClient()
	client.OnConnecting = func(address string) {
		t.Logf("connecting %s", address)
	}
	client.OnConnectFailed = func(err error) {
		t.Error(err)
	}
	loader, _ := client.Dump("127.0.0.1:6300")
	file, _ := os.Create("jacoco.exec")
	loader.Save(file)
}
