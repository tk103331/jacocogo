package analysis

const SOURCE_UNKNOWN_LINE = -1

type ElementType int

const (
	ELEMENT_METHOD ElementType = iota
	ELEMENT_CLASS
	ELEMENT_SOURCE
	ELEMENT_PACKAGE
	ELEMENT_BUNDLE
	ELEMENT_GROUP
)

type CounterEntity int

const (
	COUNTER_INSTRUCTION CounterEntity = iota
	COUNTER_BRANCH
	COUNTER_LINE
	COUNTER_COMPLEXITY
	COUNTER_METHOD
	COUNTER_CLASS
)

type CounterValue int

const (
	COUNTER_VALUE_TOTAL_COUNT CounterValue = iota
	COUNTER_VALUE_MISSED_COUNT
	COUNTER_VALUE_COVERED_COUNT
	COUNTER_VALUE_MISSED_RATIO
	COUNTER_VALUE_COVERED_RATIO
)

type CoverageStatus int

const (
	COVERAGE_EMPTY          CoverageStatus = 0x00
	COVERAGE_NOT_COVERED    CoverageStatus = 0x01
	COVERAGE_PARTLY_COVERED CoverageStatus = COVERAGE_NOT_COVERED | COVERAGE_FULLY_COVERED
	COVERAGE_FULLY_COVERED  CoverageStatus = 0x02
)

const COUNTER_SINGLETON_LIMIT = 30
const LINE_SINGLETON_INS_LIMIT = 8
const LINE_SINGLETON_BRA_LIMIT = 4

var COUNTER_SINGLETONS [][]CounterImpl
var LINE_SINGLETONS [][][][]LineImpl
var COUNTER_0_0 CounterImpl
var COUNTER_1_0 CounterImpl
var COUNTER_0_1 CounterImpl

func init() {
	COUNTER_SINGLETONS = make([][]CounterImpl, COUNTER_SINGLETON_LIMIT+1)
	for i := 0; i <= COUNTER_SINGLETON_LIMIT; i++ {
		COUNTER_SINGLETONS[i] = make([]CounterImpl, COUNTER_SINGLETON_LIMIT+1)
		for j := 0; j < COUNTER_SINGLETON_LIMIT; j++ {
			COUNTER_SINGLETONS[i][j] = CounterImpl{missed: i, covered: j}
		}
	}

	COUNTER_0_0 = COUNTER_SINGLETONS[0][0]
	COUNTER_1_0 = COUNTER_SINGLETONS[1][0]
	COUNTER_0_1 = COUNTER_SINGLETONS[0][1]

	LINE_SINGLETONS = make([][][][]LineImpl, LINE_SINGLETON_INS_LIMIT+1)
	for i := 0; i <= LINE_SINGLETON_INS_LIMIT; i++ {
		LINE_SINGLETONS[i] = make([][][]LineImpl, LINE_SINGLETON_INS_LIMIT+1)
		for j := 0; j <= LINE_SINGLETON_INS_LIMIT; j++ {
			LINE_SINGLETONS[i][j] = make([][]LineImpl, LINE_SINGLETON_BRA_LIMIT+1)
			for k := 0; k < LINE_SINGLETON_BRA_LIMIT; k++ {
				LINE_SINGLETONS[i][j][k] = make([]LineImpl, LINE_SINGLETON_BRA_LIMIT+1)
				for l := 0; l <= LINE_SINGLETON_BRA_LIMIT; l++ {
					LINE_SINGLETONS[i][j][k][l] = LineImpl{instructions: CounterImpl{missed: i, covered: j}, branches: CounterImpl{missed: k, covered: l}}
				}
			}

		}
	}

}
