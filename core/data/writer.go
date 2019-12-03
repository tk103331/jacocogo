package data

import (
	"github.com/tk103331/jacocogo/core/common"
	"io"
)

type ExecutionDataWriter struct {
	dw *common.DataWriter
}

func NewWriter(writer io.Writer) *ExecutionDataWriter {
	dataWriter := &ExecutionDataWriter{dw: common.NewWriter(writer)}
	dataWriter.writeHeader()
	return dataWriter
}

func (w *ExecutionDataWriter) Flush() error {
	return w.dw.Flush()
}
func (w *ExecutionDataWriter) writeHeader() error {
	err := w.dw.WriteByte(BLOCK_HEADER)
	if err != nil {
		return err
	}
	err = w.dw.WriteChar(MAGIC_NUMBER)
	if err != nil {
		return err
	}
	err = w.dw.WriteChar(FORMAT_VERSION)
	if err != nil {
		return err
	}
	err = w.dw.Flush()
	if err != nil {
		return err
	}
	return nil
}

func (w *ExecutionDataWriter) writeSessionInfo(info SessionInfo) error {
	err := w.dw.WriteByte(BLOCK_SESSIONINFO)
	if err != nil {
		return err
	}
	err = w.dw.WriteUTF(info.Id)
	if err != nil {
		return err
	}
	err = w.dw.WriteLong(info.Start)
	if err != nil {
		return err
	}
	err = w.dw.WriteLong(info.Dump)
	if err != nil {
		return err
	}
	err = w.dw.Flush()
	if err != nil {
		return err
	}
	return nil
}

func (w *ExecutionDataWriter) writeExecutionData(data ExecutionData) error {
	err := w.dw.WriteByte(BLOCK_EXECUTIONDATA)
	if err != nil {
		return err
	}
	err = w.dw.WriteLong(data.Id)
	if err != nil {
		return err
	}
	err = w.dw.WriteUTF(data.Name)
	if err != nil {
		return err
	}
	err = w.dw.WriteBoolArray(data.Probes)
	if err != nil {
		return err
	}
	err = w.dw.Flush()
	if err != nil {
		return err
	}
	return nil
}

func (w *ExecutionDataWriter) VisitSessionInfo(info SessionInfo) error {
	return w.writeSessionInfo(info)
}

func (w *ExecutionDataWriter) VisitExecutionData(data ExecutionData) error {
	return w.writeExecutionData(data)
}

func GetFileHeader() ([]byte, error) {
	buffer := common.NewBuffer(32)
	NewWriter(buffer)
	return buffer.Data(), nil
}
