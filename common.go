package raft

import (
	"errors"
)

type RaftState int

const (
	NODE_STATE_FOLLOWER RaftState = iota
	NODE_STATE_LEADER
	NODE_STATE_CANDIDATE
	NODE_STATE_STOP
	NODE_STATE_SNAPSHOT
)

const (
	STATE_FOLLOWER_DESCRIBE  = "follower"
	STATE_LEADER_DESCRIBE    = "leader"
	STATE_CANDIDATE_DESCRIBE = "candidate"
	STATE_STOP_DESCRIBE      = "stop"
	STATE_UNKNOWN_DESCRIBE   = "unknown"
	STATE_SNAPSHOT_DESCRIBE  = "snapshot"
)

func (r RaftState) String() string {
	switch r {
	case NODE_STATE_FOLLOWER:
		return STATE_FOLLOWER_DESCRIBE
	case NODE_STATE_LEADER:
		return STATE_LEADER_DESCRIBE
	case NODE_STATE_CANDIDATE:
		return STATE_CANDIDATE_DESCRIBE
	case NODE_STATE_STOP:
		return STATE_STOP_DESCRIBE
	case NODE_STATE_SNAPSHOT:
		return STATE_SNAPSHOT_DESCRIBE
	}
	return STATE_UNKNOWN_DESCRIBE
}

var (
	ERROR_ElectionTimeout  = errors.New("raft leader election failed")
	ERROR_HeartbeatTimeout = errors.New("raft heart beat failed")
)

type Term struct {
	StartTime      int64 `json:"start_time,omitempty"`
	Endtime        int64 `json:"endtime,omitempty"`
	Term           int64 `json:"term,omitempty"`
	Able           bool  `json:"able,omitempty"`
	EndPointNumber int   `json:"end_point_number,omitempty"`
	VoteNumber     int64 `json:"vote_number,omitempty"`
}

type OPERATION_TYPE int

const (
	OPERATION_HEARTBEAT OPERATION_TYPE = iota
	OPERATION_ELECTION
	OPERATION_APPENDLOG
	OPERATION_SNAPSHOT
	OPERATION_VOTEREQUEST
)

type Vote struct {
	CommitedID int64  `json:"commited_id,omitempty"`
	TermID     int64  `json:"term_id,omitempty"`
	NodeID     int64  `json:"node_id,omitempty"`
	Peername   string `json:"peername,omitempty"`
	TimeStamp  int64  `json:"time_stamp,omitempty"`
}
