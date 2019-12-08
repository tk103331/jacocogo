package analysis

type SourceFileCoverage interface {
	SourceNode
	PackageName() string
}

type SourceFileCoverageImpl struct {
	SourceNodeImpl
	packageName string
}

func (s SourceFileCoverageImpl) PackageName() string {
	return s.packageName
}
