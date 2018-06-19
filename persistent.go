package raft

type Persistent interface {
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
}

func (fi *FilePersistent) CreateDir(path string) error {
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
