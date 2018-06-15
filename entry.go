package raft

type Entry struct {
	logentry []byte
	Term     int64
}
