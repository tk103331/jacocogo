package data

import (
	"fmt"
	"github.com/tk103331/jacocogo/core/data/internal/data"
	"math/rand"
	"testing"
)

type _sessionVisitor struct {
}

func (_sessionVisitor) visitSessionInfo(info SessionInfo) error {
	fmt.Println(info)
	return nil
}

type _executionVisitor struct {
}

func (_executionVisitor) visitExecutionData(data ExecutionData) error {
	fmt.Println(data)
	return nil
}

func TestEmpty(t *testing.T) {
	buffer := data.NewBuffer(32)

	reader := NewReader(buffer)
	reader.ExecutionVisitor = _executionVisitor{}
	reader.SessionVisitor = _sessionVisitor{}

	_, err := reader.Read()
	assertError(t, err)
}

func TestGetFileHeader(t *testing.T) {
	header, err := GetFileHeader()

	assertNoErr(t, err)

	magic := MAGIC_NUMBER
	version := FORMAT_VERSION
	assertEqual(t, byte(1), header[0])
	assertEqual(t, magic, uint16(header[1])<<8|uint16(header[2]))
	assertEqual(t, version, uint16(header[3])<<8|uint16(header[4]))
}

func TestMultipleHeaders(t *testing.T) {
	buffer := data.NewBuffer(128)

	NewWriter(buffer)
	NewWriter(buffer)
	NewWriter(buffer)

	NewReader(buffer).Read()
}

func TestInvalidMagicNumber(t *testing.T) {
	buffer := data.NewDataBuffer(32)
	buffer.WriteByte(BLOCK_HEADER)
	buffer.WriteByte(0x12)
	buffer.WriteByte(0x34)
	buffer.Flush()

	_, err := NewReader(buffer).Read()
	assertEqual(t, InvalidExecutionDataError, err)
}

func TestInvalidVersion(t *testing.T) {
	buffer := data.NewDataBuffer(32)
	buffer.WriteByte(BLOCK_HEADER)
	buffer.WriteChar(MAGIC_NUMBER)
	version := FORMAT_VERSION - 1
	buffer.WriteChar(version)
	buffer.Flush()

	_, err := NewReader(buffer).Read()
	assertEqual(t, InvalidExecutionDataError, err)
}

func TestMissingHeader(t *testing.T) {
	buffer := data.NewDataBuffer(32)
	writer := NewWriter(buffer)
	writer.visitExecutionData(ExecutionData{Id: 0x100000000000000, Name: "Sample", Probes: createProbes(8)})
	_, err := NewReader(buffer).Read()
	assertEqual(t, InvalidExecutionDataError, err)
}
func TestUnknownBlock(t *testing.T) {
	buffer := data.NewDataBuffer(32)
	buffer.WriteByte(0xff)
	buffer.Flush()
	_, err := NewReader(buffer).Read()
	assertEqual(t, InvalidExecutionDataError, err)
}

func TestNoSessionInfoVisitor(t *testing.T) {
	buffer := data.NewDataBuffer(1024)

	NewWriter(buffer).visitSessionInfo(SessionInfo{Id: "x", Start: 0, Dump: 1})
	_, err := NewReader(buffer).Read()
	assertEqual(t, NoSessionVisitorError, err)
}

func assertEqual(t *testing.T, a, b interface{}) {
	t.Helper()
	if a != b {
		t.Errorf("Not Equal. %d %d", a, b)
	}
}

func assertTrue(t *testing.T, value bool) {
	t.Helper()
	if value {
		t.Errorf("Not True.")
	}
}
func assertFalse(t *testing.T, value bool) {
	t.Helper()
	if value {
		t.Errorf("Not False.")
	}
}

func assertError(t *testing.T, e error) {
	t.Helper()
	if e == nil {
		t.Errorf("No Error : %s", e.Error())
	}
}
func assertNoErr(t *testing.T, e error) {
	t.Helper()
	if e != nil {
		t.Errorf("Error : %s", e.Error())
	}
}

func createProbes(count uint) []bool {
	probes := make([]bool, count)
	for i := 0; i < int(count); i++ {
		probes[i] = rand.Int()/2 == 0
	}
	return probes
}
