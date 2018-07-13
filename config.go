package raft

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Address       string  `json:"address,omitempty"`
	ClientAddress string  `json:"peer,omitempty"`
	Peerlist      []*Peer `json:"peerlist,omitempty"`
	Name          string  `json:"name,omitempty"`
	ID            uint64  `json:"id,omitempty"`

	//vote timeout unit microseconds
	VoteTimeout int64 `json:"vote_timeout,omitempty"`
	//heartbeat timeout unit microseconds
	HeartbeatTimeout int64 `json:"heartbeat_timeout,omitempty"`

	//log dir
	LogDirPath   string `json:"log_dir_path,omitempty"`
	SnapshotPath string `json:"snapshot_path,omitempty"`
	SnapGap      uint64 `json:"snap_gap,omitempty"`

	//sysinfo
	RunDir string `json:"run_dir,omitempty"`
}

func NewConfigFromByte(data []byte) (*Config, error) {
	c := new(Config)
	err := json.Unmarshal(data, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func NewConfigFromFile(filepath string) (*Config, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	c := new(Config)
	err = json.Unmarshal(data, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Config) Sysinfo() []byte {
	data, err := json.Marshal(c)
	if err != nil {
		return nil
	}
	return data
}

func (c *Config) GetCurrentPath() (error, string) {
	if c.RunDir != "" {
		return nil, c.RunDir
	}
	dirPath, err := os.Getwd()
	if err != nil {
		return err, ""
	}
	c.RunDir = dirPath
	return nil, dirPath
}
