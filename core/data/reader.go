package data

import (
	"github.com/tk103331/jacocogo/core/common"
	"io"
)

// ExecutionDataReader is deserialization of execution data from binary streams.
type ExecutionDataReader struct {
	dr               *common.DataReader
	firstBlock       bool
	sessionVisitor   SessionInfoVisitor
	executionVisitor ExecutionDataVisitor
	blockVisitor     DataBlockVisitor
}

func NewReader(reader io.Reader) *ExecutionDataReader {
	return &ExecutionDataReader{dr: common.NewReader(reader)}
}

func (r *ExecutionDataReader) SetSessionVisitor(visitor SessionInfoVisitor) {
	r.sessionVisitor = visitor
}
func (r *ExecutionDataReader) SetExecutionVisitor(visitor ExecutionDataVisitor) {
	r.executionVisitor = visitor
}
func (r *ExecutionDataReader) SetDataBlockVisitor(visitor DataBlockVisitor) {
	r.blockVisitor = visitor
}

func (r *ExecutionDataReader) Read() (bool, error) {
	for {
		blockType, err := r.dr.ReadByte()
		if err != nil {
			if err == io.EOF || err == io.ErrNoProgress {
				return true, nil // EOF
			}
			return false, err
		}
		if r.firstBlock && blockType != BLOCK_HEADER {
			return false, InvalidExecutionDataError
		}
		r.firstBlock = false
		more, err := r.readBlock(blockType)
		if err != nil {
			return false, err
		}
		if !more {
			break
		}
	}
	return true, nil
}

func (r *ExecutionDataReader) readBlock(blockType byte) (bool, error) {
	switch blockType {
	case BLOCK_HEADER:
		err := r.readHeader()
		if err != nil {
			return false, err
		}
		return true, nil
	case BLOCK_SESSIONINFO:
		err := r.readSessionInfo()
		if err != nil {
			return false, err
		}
		return true, nil
	case BLOCK_EXECUTIONDATA:
		err := r.readExecutionData()
		if err != nil {
			return false, err
		}
		return true, nil
	default:
		if r.blockVisitor != nil {
			return r.blockVisitor.VisitDataBlock(DataBlock{Type: blockType})
		}
		return false, UnknownBlockTypeError
	}
}

func (r *ExecutionDataReader) readHeader() error {
	char, err := r.dr.ReadChar()
	if err != nil {
		return err
	}
	if char != MAGIC_NUMBER {
		return InvalidExecutionDataError
	}
	version, err := r.dr.ReadChar()
	if err != nil {
		return err
	}
	if version != FORMAT_VERSION {
		return InvalidExecutionDataError
	}
	return nil
}

func (r *ExecutionDataReader) readSessionInfo() error {
	if r.sessionVisitor == nil {
		return NoSessionVisitorError
	}
	id, err := r.dr.ReadUTF()
	if err != nil {
		return err
	}
	start, err := r.dr.ReadLong()
	if err != nil {
		return err
	}
	dump, err := r.dr.ReadLong()
	if err != nil {
		return err
	}
	return r.sessionVisitor.VisitSessionInfo(SessionInfo{id, start, dump})
}

func (r *ExecutionDataReader) readExecutionData() error {
	if r.executionVisitor == nil {
		return NoExecutionVisitorError
	}
	id, err := r.dr.ReadLong()
	if err != nil {
		return err
	}
	name, err := r.dr.ReadUTF()
	if err != nil {
		return err
	}
	probes, err := r.dr.ReadBoolArray()
	if err != nil {
		return err
	}
	return r.executionVisitor.VisitExecutionData(ExecutionData{id, name, probes})
}
