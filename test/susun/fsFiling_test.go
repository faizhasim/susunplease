package susun_test

import (
	"github.com/faizhasim/susunplease/internal/model"
	"github.com/faizhasim/susunplease/internal/module"
	"github.com/faizhasim/susunplease/internal/sideEffect"
	"github.com/stretchr/testify/mock"
	"path"
	"testing"
	"time"
)

type MyMockedObject struct {
	mock.Mock
}

func (m *MyMockedObject) Rename(destDir, fqSrcFile, srcFile string) error {
	m.Called(destDir, fqSrcFile, srcFile)
	//return args.Error(3)
	return nil
}

func (m *MyMockedObject) RandStringBytesMask() string {
	m.Called()
	return "random"
}

const destRootDir = "/tmp/does/not/exist"
const fqSrcFile = "somefile.pdf"

func fixture(rule *model.Rule) (*MyMockedObject, *module.SusunPdf, string) {
	mockedObject := &MyMockedObject{}
	m := &module.SusunPdf{FsFiling: sideEffect.NewFsFiling(destRootDir, mockedObject, mockedObject)}
	destDir := path.Join(destRootDir, rule.TargetDir)
	mockedObject.On("RandStringBytesMask").Return("random")
	return mockedObject, m, destDir
}

func TestFileNamingWithTimeFound(t *testing.T) {
	rule := &model.Rule{DocumentType: "docType", TargetDir: "target/dir"}
	mockedObject, m, destDir := fixture(rule)
	tm := time.Now()
	srcFile := tm.Format("2006-01-02") + " docType random.pdf"
	mockedObject.On("Rename", destDir, fqSrcFile, srcFile).Return(nil)

	err := m.FsFiling.MoveFile(&tm, fqSrcFile, rule)
	if err != nil {
		t.Error(err)
	}

	mockedObject.AssertExpectations(t)
}

func TestFileNamingWithTimeNotFound(t *testing.T) {
	rule := &model.Rule{DocumentType: "docType", TargetDir: "target/dir"}
	mockedObject, m, destDir := fixture(rule)

	srcFile := "unknowndate somefile docType random.pdf"
	mockedObject.On("Rename", destDir, fqSrcFile, srcFile).Return(nil)

	err := m.FsFiling.MoveFile(nil, fqSrcFile, rule)
	if err != nil {
		t.Error(err)
	}

	mockedObject.AssertExpectations(t)
}
