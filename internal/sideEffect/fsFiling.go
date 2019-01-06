package sideEffect

import (
	"fmt"
	"github.com/faizhasim/susunplease/internal/model"
	"path"
	"strings"
	"time"
)

type FsFiling interface {
	MoveFile(tm *time.Time, fqSrcFile string, rule *model.Rule) error
}

type fsFiling struct {
	destRootDir string
	ioOperation FsOperation
	randomGen   RandomGen
}

func (fileIo *fsFiling) MoveFile(tm *time.Time, fqSrcFile string, rule *model.Rule) error {
	_, srcFile := path.Split(fqSrcFile)
	if tm != nil {
		srcFile = tm.Format("2006-01-02") + " " + rule.DocumentType + " " + fileIo.randomGen.RandStringBytesMask() + path.Ext(srcFile)
	} else {
		srcFile = "unknowndate " + strings.TrimSuffix(srcFile, path.Ext(srcFile)) + " " + rule.DocumentType + " " + fileIo.randomGen.RandStringBytesMask() + path.Ext(srcFile)
	}

	destDir := path.Join(fileIo.destRootDir, rule.TargetDir)
	fmt.Println(destDir, ",", fqSrcFile, ",", srcFile)
	return fileIo.ioOperation.Rename(destDir, fqSrcFile, srcFile)
}

func NewFsFiling(destRootDir string, ioOperation FsOperation, randomGen RandomGen) FsFiling {
	return &fsFiling{destRootDir: destRootDir, ioOperation: ioOperation, randomGen: randomGen}
}
