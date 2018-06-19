package raft

import (
	"encoding/json"
)

type Config struct {
	Address  string  `json:"address,omitempty"`
	Peerlist []*Peer `json:"peerlist,omitempty"`
	Name     string  `json:"name,omitempty"`
	ID       uint64  `json:"id,omitempty"`

	//vote timeout unit microseconds
	VoteTimeout int64 `json:"vote_timeout,omitempty"`
	//heartbeat timeout unit microseconds
	HeartbeatTimeout int64 `json:"heartbeat_timeout,omitempty"`

	//log dir
	LogDirPath   string `json:"log_dir_path,omitempty"`
	SnapshotPath string `json:"snapshot_path,omitempty"`
	SnapGap      uint64 `json:"snap_gap,omitempty"`

	//sysinfo
	GoVersion string `json:"go_version,omitempty"`
	BuildTime string `json:"build_time,omitempty"`
	RunDir    string `json:"run_dir,omitempty"`
}

func (c *Config) Sysinfo() []byte {
	data, err := json.Marshal(c)
	if err != nil {
		return nil
	}
	return data
}
