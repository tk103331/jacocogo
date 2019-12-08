package analysis

type Line interface {
	InstructionCounter() Counter
	BranchCounter() Counter
	Status() CoverageStatus
}

type LineImpl struct {
	instructions CounterImpl
	branches     CounterImpl
}

func (l LineImpl) InstructionCounter() Counter {
	return l.instructions
}

func (l LineImpl) BranchCounter() Counter {
	return l.branches
}

func (l LineImpl) Status() CoverageStatus {
	return l.instructions.Status() | l.branches.Status()
}

func (l LineImpl) Increment(instructionCounter Counter, branchCounter Counter) LineImpl {
	return LineImpl{instructions: l.instructions.Increment(instructionCounter), branches: l.branches.Increment(branchCounter)}
}
