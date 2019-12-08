package analysis

type BundleCoverage interface {
	CoverageNode
	Packages() []PackageCoverage
}

type BundleCoverageImpl struct {
	CoverageNodeImpl
	packages []PackageCoverage
}

func (b BundleCoverageImpl) Packages() []PackageCoverage {
	return b.packages
}
