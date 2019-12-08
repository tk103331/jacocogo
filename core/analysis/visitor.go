package analysis

type CoverageVisitor interface {
	VisitCoverage(coverage ClassCoverage)
}

type AnalyzingVisitor struct {
}
