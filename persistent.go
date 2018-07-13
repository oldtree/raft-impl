package raft

import (
	"errors"
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	ErrNotDir = errors.New("path is a file ,not dir")
)

const (
	Permission uint32 = 766
)

type Persistent interface {
	Init() error
	CreateDir(path string) error
	RemoveDir(path string) error
	CreateFile(filename string) error
	RemoveFile(filename string) error
	Write([]byte) (int, error)
	Read([]byte) (int, error)
	LoadSnapshot(snapFilePath string) error
	MakeSnapshot(snapFilePath string) error
}

type FilePersistent struct {
	StoreDir string
}

func (fi *FilePersistent) Init() error {
	finfo, err := os.Stat(fi.StoreDir)
	if err != nil {
		if err == os.ErrNotExist {
			log.Infof("file [%s] not exist,could create it")
		} else {
			log.Errorf("stat file [%s] error : [%s] ", fi.StoreDir, err.Error())
			return err
		}
	}
	if finfo.IsDir() == false {
		log.Errorf("file [%s] is not dir", fi.StoreDir)
		return ErrNotDir
	}
	err = os.MkdirAll(fi.StoreDir, os.FileMode(Permission))
	if err != nil {
		log.Errorf("create dir [%s] failed")
		return err
	}
	log.Infof("create dir [%s] success ", fi.StoreDir)
	return nil
}

func (fi *FilePersistent) CreateDir(path string) error {
	err := os.MkdirAll(fi.StoreDir, os.FileMode(Permission))
	if err != nil {
		log.Errorf("create dir [%s] failed ", err.Error())
		return err
	}
	return nil
}

func (fi *FilePersistent) RemoveDir(path string) error {
	return nil
}

func (fi *FilePersistent) CreateFile(filename string) error {
	return nil
}

func (fi *FilePersistent) RemoveFile(filename string) error {
	return nil
}

func (fi *FilePersistent) Wirte([]byte) (int, error) {
	return 0, nil
}

func (fi *FilePersistent) Read([]byte) (int, error) {
	return 0, nil
}

type BoltDbPersisten struct {
}

func Init() error {
	return nil
}

func (bol *BoltDbPersisten) CreateDir(path string) error {
	return nil
}

func (bol *BoltDbPersisten) RemoveDir(path string) error {
	return nil
}

func (bol *BoltDbPersisten) CreateFile(filename string) error {
	return nil
}

func (bol *BoltDbPersisten) RemoveFile(filename string) error {
	return nil
}

func (bol *BoltDbPersisten) Wirte([]byte) (int, error) {
	return 0, nil
}

func (bol *BoltDbPersisten) Read([]byte) (int, error) {
	return 0, nil
}
