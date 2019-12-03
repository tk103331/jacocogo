package data

import (
	"github.com/tk103331/jacocogo/core/data/internal/data"
	"io"
)

type ExecutionDataWriter struct {
	dw *data.DataWriter
}

func NewWriter(writer io.Writer) *ExecutionDataWriter {
	dataWriter := &ExecutionDataWriter{dw: data.NewWriter(writer)}
	dataWriter.writeHeader()
	return dataWriter
}

func (w *ExecutionDataWriter) Flush() error {
	return w.dw.Flush()
}
func (w *ExecutionDataWriter) writeHeader() error {
	w.dw.WriteByte(BLOCK_HEADER)
	w.dw.WriteChar(MAGIC_NUMBER)
	w.dw.WriteChar(FORMAT_VERSION)
	w.dw.Flush()
	return nil
}

func (w *ExecutionDataWriter) writeSessionInfo(info SessionInfo) error {
	w.dw.WriteByte(BLOCK_SESSIONINFO)
	w.dw.WriteUTF(info.Id)
	w.dw.WriteLong(info.Start)
	w.dw.WriteLong(info.Dump)
	w.dw.Flush()
	return nil
}

func (w *ExecutionDataWriter) writeExecutionData(data ExecutionData) error {
	w.dw.WriteByte(BLOCK_EXECUTIONDATA)
	w.dw.WriteLong(data.Id)
	w.dw.WriteUTF(data.Name)
	w.dw.WriteBoolArray(data.Probes)
	w.dw.Flush()
	return nil
}

func (w *ExecutionDataWriter) visitSessionInfo(info SessionInfo) error {
	return w.writeSessionInfo(info)
}

func (w *ExecutionDataWriter) visitExecutionData(data ExecutionData) error {
	return w.writeExecutionData(data)
}

func GetFileHeader() ([]byte, error) {
	buffer := data.NewBuffer(32)
	NewWriter(buffer)
	return buffer.Data(), nil
}
