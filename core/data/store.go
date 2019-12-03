package data

// SessionInfoStore is a container to collect and merge session objects.
type SessionInfoStore struct {
	infos []SessionInfo
}

func NewSessionStore() *SessionInfoStore {
	infos := make([]SessionInfo, 0)
	return &SessionInfoStore{infos}
}

func (ss *SessionInfoStore) Infos() []SessionInfo {
	return ss.infos
}

// Writes all contained SessionInfo objects into the given visitor.
func (ss *SessionInfoStore) Accept(visitor SessionInfoVisitor) error {
	if visitor == nil {
		return NoSessionVisitorError
	}
	for _, info := range ss.infos {
		err := visitor.VisitSessionInfo(info)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ss *SessionInfoStore) VisitSessionInfo(info SessionInfo) error {
	ss.infos = append(ss.infos, info)
	return nil
}

// ExecutionDataStore is in-memory data store for execution data.
type ExecutionDataStore struct {
	entries map[int64]ExecutionData
	names   map[string]bool
}

func NewExecutionStore() *ExecutionDataStore {
	entries := make(map[int64]ExecutionData, 0)
	names := make(map[string]bool, 0)
	return &ExecutionDataStore{entries: entries, names: names}
}

func (es *ExecutionDataStore) Contents() []ExecutionData {
	contents := make([]ExecutionData, 0)
	for _, data := range es.entries {
		contents = append(contents, data)
	}
	return contents
}

func (es *ExecutionDataStore) Accept(visitor ExecutionDataVisitor) error {
	if visitor == nil {
		return NoExecutionVisitorError
	}
	for _, data := range es.entries {
		err := visitor.VisitExecutionData(data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (es *ExecutionDataStore) VisitExecutionData(data ExecutionData) error {
	es.entries[data.Id] = data
	es.names[data.Name] = true
	return nil
}
