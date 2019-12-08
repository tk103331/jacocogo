package analysis

type SourceNode interface {
	CoverageNode
	FirstLine() int
	LastLine() int
	GetLine(lineNum int) Line
}

type SourceNodeImpl struct {
	CoverageNodeImpl
	lines  []LineImpl
	offset int
}

func (s SourceNodeImpl) FirstLine() int {
	return s.offset
}

func (s SourceNodeImpl) LastLine() int {
	if len(s.lines) == 0 {
		return SOURCE_UNKNOWN_LINE
	} else {
		return s.offset + len(s.lines) - 1
	}
}

func (s SourceNodeImpl) GetLine(lineNum int) Line {
	if len(s.lines) == 0 || lineNum < s.FirstLine() || lineNum > s.LastLine() {
		return nil
	}
	return s.lines[lineNum]
}

func (s SourceNodeImpl) Increment(node SourceNode) {
	s.instructionCounter = s.instructionCounter.Increment(node.InstructionCounter())
	s.lineCounter = s.lineCounter.Increment(node.LineCounter())
	s.branchCounter = s.branchCounter.Increment(node.BranchCounter())
	s.methodCounter = s.methodCounter.Increment(node.MethodCounter())
	s.complexityCounter = s.complexityCounter.Increment(node.ComplexityCounter())
	s.classCounter = s.classCounter.Increment(node.ClassCounter())

	firstLine := node.FirstLine()
	if firstLine != SOURCE_UNKNOWN_LINE {
		lastLine := node.LastLine()
		for i := firstLine; i <= lastLine; i++ {
			line := node.GetLine(i)
			s.incrLine(line.InstructionCounter(), line.BranchCounter(), i)
		}
	}
}

func (s SourceNodeImpl) incrLine(instructionCounter Counter, branchCounter Counter, lineNum int) {
	line := s.GetLine(lineNum).(LineImpl)
	oldTotal := line.InstructionCounter().TotalCount()
	oldCovered := line.InstructionCounter().CoveredCount()
	s.lines[lineNum-s.offset] = line.Increment(instructionCounter, branchCounter)
	if instructionCounter.TotalCount() > 0 {
		if instructionCounter.CoveredCount() == 0 {
			if oldTotal == 0 {
				s.lineCounter = s.lineCounter.Increment(COUNTER_1_0)
			}
		} else {
			if oldTotal == 0 {
				s.lineCounter = s.lineCounter.Increment(COUNTER_0_1)
			} else if oldCovered == 0 {
				s.lineCounter = s.lineCounter.Increment(CounterImpl{missed: -1, covered: 1})
			}
		}
	}
}

func (s SourceNodeImpl) IncrementLine(instructionCounter Counter, branchCounter Counter, lineNum int) {
	if lineNum != SOURCE_UNKNOWN_LINE {
		s.incrLine(instructionCounter, branchCounter, lineNum)
	}
	s.instructionCounter = s.instructionCounter.Increment(instructionCounter)
	s.branchCounter = s.branchCounter.Increment(branchCounter)
}
