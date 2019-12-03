package data

type SessionInfoVisitor interface {
	VisitSessionInfo(info SessionInfo) error
}

type ExecutionDataVisitor interface {
	VisitExecutionData(data ExecutionData) error
}

type DataBlockVisitor interface {
	VisitDataBlock(block DataBlock) (bool, error)
}
