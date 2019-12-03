package data

import (
	"github.com/tk103331/jacocogo/core/data/internal/data"
	"io"
)

// ExecutionDataReader is deserialization of execution data from binary streams.
type ExecutionDataReader struct {
	dr               *data.DataReader
	firstBlock       bool
	SessionVisitor   SessionInfoVisitor
	ExecutionVisitor ExecutionDataVisitor
}

func NewReader(reader io.Reader) *ExecutionDataReader {
	return &ExecutionDataReader{dr: data.NewReader(reader)}
}

func (r *ExecutionDataReader) Read() (bool, error) {
	for {
		blockType, err := r.dr.ReadByte()
		if err != nil {
			return false, err // EOF
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
	if r.SessionVisitor == nil {
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
	return r.SessionVisitor.visitSessionInfo(SessionInfo{id, start, dump})
}

func (r *ExecutionDataReader) readExecutionData() error {
	if r.ExecutionVisitor == nil {
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
	return r.ExecutionVisitor.visitExecutionData(ExecutionData{id, name, probes})
}
