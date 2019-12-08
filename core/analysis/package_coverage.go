package analysis

type PackageCoverage interface {
	CoverageNode
	Classes() []ClassCoverage
	SourceFiles() []SourceFileCoverage
}

type PackageCoverageImpl struct {
	CoverageNodeImpl
	classes     []ClassCoverage
	sourceFiles []SourceFileCoverage
}

func (p PackageCoverageImpl) Classes() []ClassCoverage {
	return p.classes
}

func (p PackageCoverageImpl) SourceFiles() []SourceFileCoverage {
	return p.sourceFiles
}
