package raft

import (
	"encoding/json"
)

type Config struct {
	Address  string
	Peerlist []*Peer
	Name     string
	ID       uint64

	//vote timeout unit microseconds
	VoteTimeout int64
	//heartbeat timeout unit microseconds
	HeartbeatTimeout int64

	//log dir
	LogDirPath   string
	SnapshotPath string
	SnapGap      uint64

	//sysinfo
	GoVersion string
	BuildTime string
	RunDir    string
}

func (c *Config) Sysinfo() []byte {
	data, err := json.Marshal(c)
	if err != nil {
		return nil
	}
	return data
}
