package runtime

import (
	"github.com/tk103331/jacocogo/core/common"
	execdata "github.com/tk103331/jacocogo/core/data"
	"io"
)

type ControlReader struct {
	*execdata.ExecutionDataReader
	r              *common.DataReader
	CommandVisitor CommandVisitor
}

func NewControlReader(reader io.Reader) *ControlReader {
	execReader := execdata.NewReader(reader)
	controlReader := &ControlReader{execReader, common.NewReader(reader), nil}
	execReader.SetDataBlockVisitor(controlReader)
	return controlReader
}

func (cr *ControlReader) VisitDataBlock(block execdata.DataBlock) (bool, error) {
	switch block.Type {
	case BLOCK_CMDDUMP:
		cr.readDumpCommand()
		return true, nil
	case BLOCK_CMDOK:
		return false, nil
	default:
		return false, execdata.UnknownBlockTypeError
	}
}

func (cr *ControlReader) readDumpCommand() error {
	if cr.CommandVisitor == nil {
		return NoCommandVisitorError
	}
	dump, err := cr.r.ReadBool()
	if err != nil {
		return err
	}
	reset, err := cr.r.ReadBool()
	if err != nil {
		return err
	}
	return cr.CommandVisitor.VisitDumpCommand(dump, reset)
}
