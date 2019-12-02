package data

import (
	"errors"
	"github.com/tk103331/jacocogo/core/data/internal/data"
	"io"
)

// ExecutionDataReader is deserialization of execution data from binary streams.
type ExecutionDataReader struct {
	dr               *data.DataReader
	firstBlock       bool
	sessionVisitor   SessionInfoVisitor
	executionVisitor ExecutionDataVisitor
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
			return false, InvalidExecutionDataError
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
		return true, nil
	default:
		return false, errors.New("unknown block type")
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
		return errors.New("incompatible execution data version")
	}
	return nil
}

func (r *ExecutionDataReader) readSessionInfo() error {
	if r.sessionVisitor == nil {
		return errors.New("no session info visitor")
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
	return r.sessionVisitor.visitSessionInfo(SessionInfo{id, start, dump})
}

func (r *ExecutionDataReader) readExecutionData() error {
	if r.executionVisitor == nil {
		return errors.New("no execution data visitor")
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
	return r.executionVisitor.visitExecutionData(ExecutionData{id, name, probes})
}
