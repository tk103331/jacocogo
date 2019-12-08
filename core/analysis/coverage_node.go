package analysis

type CoverageNode interface {
	ElementType() ElementType
	Name() string
	InstructionCounter() Counter
	BranchCounter() Counter
	LineCounter() Counter
	ComplexityCounter() Counter
	MethodCounter() Counter
	ClassCounter() Counter
	ContainsCode() bool
}

type CoverageNodeImpl struct {
	elementType        ElementType
	name               string
	instructionCounter CounterImpl
	lineCounter        CounterImpl
	branchCounter      CounterImpl
	complexityCounter  CounterImpl
	methodCounter      CounterImpl
	classCounter       CounterImpl
}

func (c CoverageNodeImpl) ElementType() ElementType {
	return c.elementType
}

func (c CoverageNodeImpl) Name() string {
	return c.name
}

func (c CoverageNodeImpl) InstructionCounter() Counter {
	return c.instructionCounter
}

func (c CoverageNodeImpl) BranchCounter() Counter {
	return c.branchCounter
}

func (c CoverageNodeImpl) LineCounter() Counter {
	return c.lineCounter
}

func (c CoverageNodeImpl) ComplexityCounter() Counter {
	return c.complexityCounter
}

func (c CoverageNodeImpl) MethodCounter() Counter {
	return c.methodCounter
}

func (c CoverageNodeImpl) ClassCounter() Counter {
	return c.classCounter
}

func (c CoverageNodeImpl) ContainsCode() bool {
	return c.instructionCounter.TotalCount() != 0
}

func (c CoverageNodeImpl) Increment(node CoverageNode) CoverageNode {
	c.instructionCounter = c.instructionCounter.Increment(node.InstructionCounter())
	c.lineCounter = c.lineCounter.Increment(node.LineCounter())
	c.branchCounter = c.branchCounter.Increment(node.BranchCounter())
	c.methodCounter = c.methodCounter.Increment(node.MethodCounter())
	c.complexityCounter = c.complexityCounter.Increment(node.ComplexityCounter())
	c.classCounter = c.classCounter.Increment(node.ClassCounter())
	return c
}
