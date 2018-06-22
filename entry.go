package raft

type LogEntry struct {
	logentry []byte
	Term     int64
}
