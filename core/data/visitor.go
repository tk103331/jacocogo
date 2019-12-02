package data

type SessionInfoVisitor interface {
	visitSessionInfo(info SessionInfo) error
}

type ExecutionDataVisitor interface {
	visitExecutionData(data ExecutionData) error
}
