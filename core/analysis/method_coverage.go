package analysis

import "math"

type MethodCoverage interface {
	SourceNode
	Description() string
	Signature() string
}

type MethodCoverageImpl struct {
	SourceNodeImpl
	description string
	signature   string
}

func (m MethodCoverageImpl) Description() string {
	return m.description
}

func (m MethodCoverageImpl) Signature() string {
	return m.signature
}

func (s MethodCoverageImpl) IncrementLine(instructionCounter Counter, branchCounter Counter, lineNum int) {
	s.SourceNodeImpl.IncrementLine(instructionCounter, branchCounter, lineNum)
	if branchCounter.TotalCount() > 1 {
		c := int(math.Max(0, float64(branchCounter.CoveredCount()-1)))
		m := int(math.Max(0, float64(branchCounter.TotalCount()-c-1)))
		s.complexityCounter = s.complexityCounter.Increment(CounterImpl{missed: m, covered: c})
	}
}

func (s MethodCoverageImpl) IncrementMethod() {
	var base Counter
	if s.instructionCounter.CoveredCount() == 0 {
		base = COUNTER_1_0
	} else {
		base = COUNTER_0_1
	}
	s.branchCounter = s.branchCounter.Increment(base)
	s.complexityCounter = s.complexityCounter.Increment(base)
}
