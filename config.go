package raft

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var (
	ErrCheckConfigFailed              = errors.New("config check failed")
	ErrCheckConfigAddressFailed       = errors.New("config address check failed")
	ErrCheckConfigClientAddressFailed = errors.New("config cleint address check failed")
	ErrCheckConfigPeerlistFailed      = errors.New("config peer list failed")
)

const (
	META               = "-meta.json"
	SNAPSHOT           = "-snapshot.log"
	PathFormat         = "/d%/%d-%d.log" //day-term-commit
	MetaPathFormat     = "/d%/%d-%d-meta.json"
	SnapshotPathFormat = "/d%/%d-%d-snapshot.log"
)

func GetFormatPath(term, commit int) (metaPath, snapPath, logPath string) {
	day := time.Now().Day()
	metaPath = fmt.Sprintf(MetaPathFormat, day, term, commit)
	snapPath = fmt.Sprintf(SnapshotPathFormat, day, term, commit)
	logPath = fmt.Sprintf(PathFormat, day, term, commit)
	return
}

type Config struct {
	Address       string  `json:"address,omitempty"`
	ClientAddress string  `json:"peer,omitempty"`
	Peerlist      []*Peer `json:"peerlist,omitempty"`
	Name          string  `json:"name,omitempty"`
	ID            uint64  `json:"id,omitempty"`

	// vote timeout unit microseconds
	VoteTimeout int64 `json:"vote_timeout,omitempty"`
	// heartbeat timeout unit microseconds
	HeartbeatTimeout int64 `json:"heartbeat_timeout,omitempty"`
	// term circle
	TermCircle int64 `json:"term_circle,omitempty"`

	//log dir
	LogDirPath   string `json:"log_dir_path,omitempty"`
	SnapshotPath string `json:"snapshot_path,omitempty"`

	//sysinfo
	RunDir string `json:"run_dir,omitempty"`
}

func (cfg *Config) VailedCheck() (ok bool, err error) {
	if len(cfg.Address) == 0 {
		ok = false
		err = ErrCheckConfigAddressFailed
		return
	}
	if len(cfg.ClientAddress) == 0 {
		ok = false
		err = ErrCheckConfigClientAddressFailed
		return
	}

	if len(cfg.Peerlist) < 3 {
		ok = false
		err = ErrCheckConfigPeerlistFailed
		return
	}

	if cfg.VoteTimeout < 0 {
		ok = false
		err = ErrCheckConfigFailed
		return
	}

	if cfg.HeartbeatTimeout < 0 {
		ok = false
		err = ErrCheckConfigFailed
		return
	}

	if cfg.LogDirPath == "" {
		cfg.LogDirPath, _ = os.Getwd()
	}

	ok, err = true, nil
	return
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
