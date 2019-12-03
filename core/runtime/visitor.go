package runtime

type CommandVisitor interface {
	VisitDumpCommand(dump, reset bool) error
}
