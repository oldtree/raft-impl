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

type BoltDbPersisten struct {
}
