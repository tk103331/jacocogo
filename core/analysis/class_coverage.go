package analysis

import "strings"

type ClassCoverage interface {
	SourceNode
	Id() int64
	Signature() string
	SuperName() string
	InterfaceNames() []string
	PackageName() string
	SourceFileName() string
	Methods() []MethodCoverage
}

type ClassCoverageImpl struct {
	SourceNodeImpl
	id             int64
	noMatch        bool
	signature      string
	superName      string
	interfaces     []string
	sourceFileName string
	methods        []MethodCoverage
}

func (c ClassCoverageImpl) Id() int64 {
	return c.id
}

func (c ClassCoverageImpl) Signature() string {
	return c.signature
}

func (c ClassCoverageImpl) SuperName() string {
	return c.superName
}

func (c ClassCoverageImpl) InterfaceNames() []string {
	return c.interfaces
}

func (c ClassCoverageImpl) PackageName() string {
	name := c.Name()
	pos := strings.LastIndex(name, "/")
	if pos == -1 {
		return ""
	} else {
		return string([]rune(name)[:pos])
	}
}

func (c ClassCoverageImpl) SourceFileName() string {
	return c.sourceFileName
}

func (c ClassCoverageImpl) Methods() []MethodCoverage {
	return c.methods
}
