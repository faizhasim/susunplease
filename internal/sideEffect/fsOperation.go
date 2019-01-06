package sideEffect

import (
	"os"
	"path"
)

type FsOperation interface {
	Rename(destDir, fqSrcFile, srcFile string) error
}

type fsOperation struct{}

func NewFsOperation() FsOperation {
	return &fsOperation{}
}

func (ioOperation *fsOperation) Rename(destDir, fqSrcFile, srcFile string) error {
	err := os.MkdirAll(destDir, os.ModePerm)
	if err != nil {
		return err
	}

	if err := os.Rename(fqSrcFile, path.Join(destDir, srcFile)); err != nil {
		return err
	}
	return nil
}
