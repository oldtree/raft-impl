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
)

const (
	STATE_FOLLOWER_DESCRIBE  = "follower"
	STATE_LEADER_DESCRIBE    = "leader"
	STATE_CANDIDATE_DESCRIBE = "candidate"
	STATE_STOP_DESCRIBE      = "stop"
	STATE_UNKNOWN_DESCRIBE   = "unknown"
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
	}
	return STATE_UNKNOWN_DESCRIBE
}

var (
	ERROR_ElectionTimeout  = errors.New("raft leader election failed")
	ERROR_HeartbeatTimeout = errors.New("raft heart beat failed")
)