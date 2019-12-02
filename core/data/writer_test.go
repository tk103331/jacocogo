package data

import (
	"fmt"
	"testing"
)

func TestGetFileHeader(t *testing.T) {
	header, err := GetFileHeader()
	if err != nil {
		t.Fatal(err)
	}
	magic := MAGIC_NUMBER
	version := FORMAT_VERSION
	t.Log(fmt.Sprintf("%d=%d", 1, header[0]))
	t.Log(fmt.Sprintf("%d=%d", magic, int(header[1])<<8|int(header[2])))
	t.Log(fmt.Sprintf("%d=%d", version, int(header[3])<<8|int(header[4])))

}
