package tools

import (
	"github.com/tk103331/jacocogo/core/data"
	"io"
)

type FileLoader struct {
	sessionStore   *data.SessionInfoStore
	executionStore *data.ExecutionDataStore
}

func NewFileLoader() *FileLoader {
	return &FileLoader{data.NewSessionStore(), data.NewExecutionStore()}
}

func (fl *FileLoader) SessionStore() *data.SessionInfoStore {
	return fl.sessionStore
}

func (fl *FileLoader) ExecutionStore() *data.ExecutionDataStore {
	return fl.executionStore
}

func (fl *FileLoader) Load(reader io.Reader) error {
	dataReader := data.NewReader(reader)
	dataReader.SetSessionVisitor(fl.sessionStore)
	dataReader.SetExecutionVisitor(fl.executionStore)
	_, err := dataReader.Read()
	return err
}

func (fl *FileLoader) Save(writer io.Writer) error {
	dataWriter := data.NewWriter(writer)
	err := fl.sessionStore.Accept(dataWriter)
	if err != nil {
		return err
	}
	return fl.executionStore.Accept(dataWriter)
}
