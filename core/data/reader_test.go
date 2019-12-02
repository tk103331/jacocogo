package data

import (
	"fmt"
	"os"
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

func TestExecutionDataReader_Read(t *testing.T) {

	file, err := os.Open("../../jacoco.exec")
	if err != nil {
		t.Error(err)
	}
	reader := NewReader(file)

	reader.sessionVisitor = _sessionVisitor{}
	reader.executionVisitor = _executionVisitor{}

	_, err = reader.Read()
	if err != nil {
		t.Error(err)
		panic(err)
	}
}
