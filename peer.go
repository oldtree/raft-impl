package raft

import (
	"context"
	"errors"
	"sync"

	_ "eco/raft/protobuffer"

	_ "github.com/sirupsen/logrus"
)

type Peer struct {
	Addr           string    `json:"addr,omitempty"`
	ID             int64     `json:"id,omitempty"`
	Name           string    `json:"name,omitempty"`
	Role           RaftState `json:"role,omitempty"`
	Term           int64     `json:"term,omitempty"`
	LastCommit     int64     `json:"last_commit,omitempty"`
	LastActiveTime int64     `json:"last_active_time,omitempty"`
}

func NewPeer(address string, id uint64, name string, role RaftState, term int64, lastCommit int64, lastActiveTimeStamp int64) *Peer {
	return &Peer{
		Addr:           address,
		ID:             int64(id),
		Name:           name,
		Role:           role,
		Term:           term,
		LastCommit:     lastCommit,
		LastActiveTime: lastActiveTimeStamp,
	}
}

func (peer *Peer) SendVoteRequest(ctx context.Context) error {

	return nil
}

func (peer *Peer) SendVoteResponse(ctx context.Context) error {

	return nil
}

func (peer *Peer) SendHeartBeatRequest(ctx context.Context) error {
	return nil
}

func (peer *Peer) SendHeartBeatResponse(ctx context.Context) error {
	return nil
}

func (peer *Peer) SendApplyLogRequest(ctx context.Context) error {
	return nil
}

func (peer *Peer) SendApplyLogResponse(ctx context.Context) error {
	return nil
}

type PeerList struct {
	Peerlist map[string]*Peer
	sync.Mutex
}

func NewPeerList() *PeerList {
	return &PeerList{
		Peerlist: make(map[string]*Peer),
	}
}

func (pl *PeerList) SendVoteRequest(ctx context.Context) error {

	return nil
}

func (pl *PeerList) SendVoteResponse(ctx context.Context) error {
	return nil
}

func (pl *PeerList) SendHeartBeatRequest(ctx context.Context) error {
	return nil
}

func (pl *PeerList) SendHeartBeatResponse(ctx context.Context) error {
	return nil
}

func (pl *PeerList) SendApplyLogRequest(ctx context.Context) error {
	return nil
}

func (pl *PeerList) SendApplyLogResponse(ctx context.Context) error {
	return nil
}

func (pl *PeerList) QuorumValue() int {
	if len(pl.Peerlist) == 1 {
		return 1
	}
	return len(pl.Peerlist)/2 + 1
}

func (pl *PeerList) AddPeer(pp *Peer) (ok bool, err error) {
	pl.Lock()
	defer pl.Unlock()
	if _, exist := pl.Peerlist[pp.Name]; exist {
		ok, err = true, nil
		return
	} else {
		pl.Peerlist[pp.Name] = pp
	}
	ok, err = true, nil
	return
}

func (pl *PeerList) RemovePeer(p *Peer) (ok bool, err error) {
	pl.Lock()
	defer pl.Unlock()
	if _, exist := pl.Peerlist[p.Name]; exist {
		delete(pl.Peerlist, p.Name)
		ok, err = true, nil
		return
	} else {
		return false, errors.New("peer is not exist")
	}
}
