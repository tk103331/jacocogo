package analysis

type Counter interface {
	TotalCount() int
	CoveredCount() int
	MissedCount() int
	CoveredRatio() float32
	MissedRatio() float32
	Status() CoverageStatus
}

type CounterImpl struct {
	missed  int
	covered int
}

func (c CounterImpl) TotalCount() int {
	return c.missed + c.covered
}

func (c CounterImpl) CoveredCount() int {
	return c.covered
}

func (c CounterImpl) MissedCount() int {
	return c.missed
}

func (c CounterImpl) CoveredRatio() float32 {
	return float32(c.covered) / float32(c.missed+c.covered)
}

func (c CounterImpl) MissedRatio() float32 {
	return float32(c.missed) / float32(c.missed+c.covered)
}

func (c CounterImpl) Status() CoverageStatus {
	var status CoverageStatus
	if c.covered > 0 {
		status = COVERAGE_FULLY_COVERED
	} else {
		status = COVERAGE_EMPTY
	}
	if c.missed > 0 {
		status |= COVERAGE_NOT_COVERED
	}
	return status
}

func (c CounterImpl) Increment(counter Counter) CounterImpl {
	return CounterImpl{c.missed + counter.MissedCount(), c.covered + counter.CoveredCount()}
}
