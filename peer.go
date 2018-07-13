package raft

import (
	"context"

	log "github.com/sirupsen/logrus"
)

type Peer struct {
	Addr           string        `json:"addr,omitempty"`
	ID             int64         `json:"id,omitempty"`
	Name           string        `json:"name,omitempty"`
	Role           RaftState     `json:"role,omitempty"`
	Term           int64         `json:"term,omitempty"`
	LastCommit     int64         `json:"last_commit,omitempty"`
	StopChan       chan struct{} `json:"-"`
	LastActiveTime int64         `json:"last_active_time,omitempty"`
}

// vote request
//The candidate then requests votes from other peer.
func (peer *Peer) SendVoteRequest(ctx context.Context) error {
	log.Infof("send vote request to : %s", peer.Addr)
	return nil
}

func (peer *Peer) SendVoteResponse(ctx context.Context, address string) error {
	log.Infof("send vote response to : [%s] ", address)
	return nil
}

func (peer *Peer) SendheartBeat() error {
	return nil
}
