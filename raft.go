package raft

import (
	"sync"

	log "github.com/sirupsen/logrus"
)

type HttpServer struct {
	Address string
}

type RpcServer struct {
	Address string
}

type CtlCommandType uint

const (
	JoinMember CtlCommandType = iota
	ReleaseSnapshot
	TransSnapshot
)

type CtlCommand struct {
	CMD    CtlCommandType
	Pramas interface{}
}

type Raft struct {
	Mutex sync.RWMutex
	Node  *Node

	ElectionTimeOut  int64
	HeartBeatTimeout int64

	CmdChan chan *CtlCommand

	StopChan chan struct{}

	HttpApiServer *HttpServer
	RpcApiServer  *RpcServer

	Peers        map[string]*Peer
	GlobalConfig *Config
}

func NewRaft(cfg *Config) *Raft {
	return &Raft{
		ElectionTimeOut:  3,
		HeartBeatTimeout: 3,
		StopChan:         make(chan struct{}, 1),
		CmdChan:          make(chan *CtlCommand, 1),
		GlobalConfig:     cfg,
		HttpApiServer:    new(HttpServer),
		RpcApiServer:     new(RpcServer),
		Peers:            make(map[string]*Peer),
	}
}

func (rf *Raft) Init() error {
	return nil
}

//This process is called Leader Election.
//All changes to the system now go through the leader.
func (rf *Raft) LeaderElection() error {
	log.Info("start leader election")
	//In Raft there are two timeout settings which control elections.
	//1:First is the election timeout.
	// The election timeout is the amount of time a follower waits until becoming a candidate.
	// The election timeout is randomized to be between 150ms and 300ms.
	// After the election timeout the follower becomes a candidate and starts a new election term ,and vote for itself first
	// and sends out Request Vote messages to other nodes.
	// If the receiving node hasn't voted yet in this term then it votes for the candidate,and the node resets its election timeout
	// Once a candidate has a majority of votes it becomes leader.The leader begins sending out Append Entries messages to its followers.
	// This election term will continue until a follower stops receiving heartbeats and becomes a candidate.
	log.Info("end leader election")
	return nil
}

func (rf *Raft) HeartBeat() error {
	//These messages are sent in intervals specified by the heartbeat timeout.
	//Followers then respond to each Append Entries message.
	return nil
}

//log
//Each change is added as an entry in the node's log.
//This log entry is currently uncommitted so it won't update the node's value.
//To commit the entry the node first replicates it to the follower nodes...
//then the leader waits u	ntil a majority of nodes have written the entry.
//The entry is now committed on the leader node and the node state is update.
//The leader then notifies the followers that the entry is committed.
//The cluster has now come to consensus about the system state.
//This process is called Log Replication.
func (rf *Raft) LogReplication() error {

	log.Infof("raft log replication start")
	//Once we have a leader elected we need to replicate all changes to our system to all nodes.
	//This is done by using the same Append Entries message that was used for heartbeats.
	//step :
	/*
		1:First a client sends a change to the leader.
		2:The change is appended to the leader's log...
		3:then the change is sent to the followers on the next heartbeat.
		4:An entry is committed once a majority of followers acknowledge it...
		5:and a response is sent to the client.
	*/

	log.Infof("raft log replication end")
	return nil
}

func (rf *Raft) AsFollower() error {
	return nil
}

func (rf *Raft) AsLeader() error {
	return nil
}

func (rf *Raft) AsCandidater() error {
	return nil
}

func (rf *Raft) loop() error {
	for {
		select {}
	}
	return nil
}
