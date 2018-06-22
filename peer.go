package raft

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
)

type Peer struct {
	Addr           string
	ID             int64
	Name           string
	Role           RaftState
	Term           int64
	LastCommit     int64
	StopChan       chan struct{}
	LastActiveTime time.Time
}

// vote request
//The candidate then requests votes from other peer.
func (peer *Peer) SendVoteRequest(ctx context.Context) error {
	log.Infof("send need vote request to : %s", peer.Addr)
	return nil
}
