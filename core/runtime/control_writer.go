package runtime

import (
	"github.com/tk103331/jacocogo/core/common"
	execdata "github.com/tk103331/jacocogo/core/data"
	"io"
)

type ControlWriter struct {
	*execdata.ExecutionDataWriter
	w *common.DataWriter
}

func NewControlWriter(writer io.Writer) *ControlWriter {
	return &ControlWriter{execdata.NewWriter(writer), common.NewWriter(writer)}
}

func (cw *ControlWriter) sendCmdOk() error {
	return cw.w.WriteByte(BLOCK_CMDOK)
}

func (cw *ControlWriter) VisitDumpCommand(dump, reset bool) error {
	cw.w.WriteByte(BLOCK_CMDDUMP)
	cw.w.WriteBool(dump)
	cw.w.WriteBool(reset)
	return nil
}
