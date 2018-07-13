package raft

import (
	"context"
	"math/rand"
	"sync"
	"time"

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
	StartRaft CtlCommandType = iota
	JoinMember
	ReleaseSnapshot
	TransSnapshot
	InstallSnapshot
)

const (
	MinHeartbeatTimeout int64 = 500  //ms
	MaxHeartbeatTimeout int64 = 2000 //ms
)

type CtlCommand struct {
	CMD    CtlCommandType
	Params interface{}
}

func NewCtlCmd(cmd CtlCommandType, params interface{}) *CtlCommand {
	return &CtlCommand{
		CMD:    cmd,
		Params: params,
	}
}

type Raft struct {
	Mutex        sync.RWMutex
	Node         *Peer
	StateMachine *StateMachine

	initOnce sync.Once

	ElectionTimeOut  int64
	HeartBeatTimeout int64

	VoteFor string
	IsVoted bool

	Heartbeat chan struct{}
	LogEntry  chan struct{}
	StopChan  chan struct{}

	CmdChan chan *CtlCommand

	HttpApiServer *HttpServer
	RpcApiServer  *RpcServer

	Peers        map[string]*Peer
	GlobalConfig *Config
}

func (rf *Raft) QuorumValue() int {
	if len(rf.Peers) == 1 {
		return 1
	}
	return len(rf.Peers)/2 + 1
}

func (rf *Raft) Isself(p *Peer) bool {
	return rf.Node.Name == p.Name
}

func NewRaft(cfg *Config) *Raft {
	return &Raft{
		ElectionTimeOut:  3000,
		StateMachine:     NewStateMachine(),
		HeartBeatTimeout: MaxHeartbeatTimeout,
		StopChan:         make(chan struct{}, 1),
		CmdChan:          make(chan *CtlCommand, 1),
		GlobalConfig:     cfg,
		HttpApiServer:    new(HttpServer),
		RpcApiServer:     new(RpcServer),
		Peers:            make(map[string]*Peer),
	}
}

func (rf *Raft) Init() error {
	rf.initOnce.Do(func() {
		log.Infof("start init raft node [%s] ", rf.GlobalConfig.Name)
	})
	rf.StateMachine.State = NODE_STATE_CANDIDATE
	rf.CmdChan <- NewCtlCmd(StartRaft, nil)
	go rf.runloop()
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
	ctx, cancelfunc := context.WithDeadline(context.Background(), time.Now().Add(time.Millisecond*time.Duration(rf.ElectionTimeOut)))
	defer cancelfunc()
	for endPointName, endpoint := range rf.Peers {
		log.Infof("send vote request to : ", endPointName)
		go endpoint.SendVoteRequest(ctx)
	}

	log.Info("end leader election")
	return nil
}

func (rf *Raft) HeartBeat() error {
	//These messages are sent in intervals specified by the heartbeat timeout.
	//Followers then respond to each Append Entries message.
	tick := time.NewTimer(time.Duration(rf.HeartBeatTimeout))
	var gap int64
	var int64rand = rand.New(rand.NewSource(time.Now().Unix()))
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			gap = int64rand.Int63n(MaxHeartbeatTimeout)
			rf.Heartbeat <- struct{}{}
			if gap < MinHeartbeatTimeout {
				gap = MinHeartbeatTimeout
			}
			tick.Reset(time.Duration(gap))
		case <-rf.StopChan:
			log.Infof("raft heart beat timer[%s] is exit ", time.Now().String())
			return nil
		}
	}
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
	for key, follwerNode := range rf.Peers {
		log.Infof("log replication to node [%s] ", key)
		if follwerNode != nil && rf.Isself(follwerNode) {
			continue
		}
	}

	log.Infof("raft log replication end")
	return nil
}

func (rf *Raft) ApplyLog() error {
	log.Infof("apply raft log")
	return nil
}

func (rf *Raft) runAsFollower() error {

	for {
		select {
		case <-rf.StopChan:
			log.Infof("raft server is close")
			rf.StateMachine.State = NODE_STATE_STOP
			break
		case <-rf.Heartbeat:
		}
	}
	return nil
}

func (rf *Raft) runAsLeader() error {
	for {
		select {
		case <-rf.StopChan:
			log.Infof("raft server is close")
			rf.StateMachine.State = NODE_STATE_STOP
			break
		}
	}
	return nil
}

func (rf *Raft) runAsCandidater() error {
	for {
		select {
		case <-rf.StopChan:
			log.Infof("raft server is close")
			rf.StateMachine.State = NODE_STATE_STOP
			break
		case <-rf.Heartbeat:
		}
	}
	return nil
}

func (rf *Raft) runloop() error {
	for {
		switch rf.StateMachine.State {
		case NODE_STATE_LEADER:
			log.Infof("node state is [%s] ", NODE_STATE_LEADER.String())
			rf.runAsLeader()
		case NODE_STATE_FOLLOWER:
			log.Infof("node state is [%s] ", NODE_STATE_FOLLOWER.String())
			rf.runAsFollower()
		case NODE_STATE_CANDIDATE:
			log.Infof("node state is [%s] ", NODE_STATE_CANDIDATE.String())
			rf.runAsCandidater()
		case NODE_STATE_STOP:
			log.Warnf("node state is [%s] ", NODE_STATE_STOP.String())
			break
		}
	}
	return nil
}
