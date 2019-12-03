package data

import "errors"

const FORMAT_VERSION uint16 = 0x1007

// Magic number in header for file format identification.
const MAGIC_NUMBER uint16 = 0xC0C0

// BLOCK_HEADER is block identifier for file headers.
const BLOCK_HEADER byte = 0x01

// BLOCK_SESSIONINFO is block identifier for session information.
const BLOCK_SESSIONINFO byte = 0x10

// BLOCK_EXECUTIONDATA is block identifier for execution data of a single class.
const BLOCK_EXECUTIONDATA byte = 0x11

var InvalidExecutionDataError error = errors.New("invalid execution data file")
var NoSessionVisitorError error = errors.New("no session info visitor")
var NoExecutionVisitorError error = errors.New("no session info visitor")
var IncompatibleVersionError error = errors.New("incompatible execution data version")
var UnknownBlockTypeError error = errors.New("unknown block type")

// SessionInfo is a session which was the source of execution data.
type SessionInfo struct {
	Id    string
	Start int64
	Dump  int64
}

// ExecutionData is execution data for a single Java class.
// While instances are immutable care has to be taken about the probe data array of type boolean[] which can be modified.
type ExecutionData struct {
	Id     int64
	Name   string
	Probes []bool
}

type DataBlock struct {
	Type byte
}
