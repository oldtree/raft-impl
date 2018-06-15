package raft

import (
	log "github.com/sirupsen/logrus"
)

type Node struct {
	NodeState RaftState
}

func NewNode() *Node {
	return &Node{
		NodeState: NODE_STATE_FOLLOWER,
	}
}

//If followers don't hear from a leader then they can become a candidate.

//The candidate becomes the leader if it gets votes from a majority of nodes.
func (node *Node) TransStatus(targetState RaftState, condition interface{}) (RaftState, error) {
	log.Infof("TransStatus from [%d] to [%d],condition : [%v] ", condition)
	if node.NodeState == targetState {
		return node.NodeState, nil
	}
	return node.NodeState, nil
}

// vote request
//The candidate then requests votes from other nodes.
func (node *Node) SendVoteRequest(addresslist []string) error {
	log.Infof("send need vote request to : %s", addresslist)
	return nil
}

//vote response
//Nodes will reply with their vote.
func (node *Node) SendVoteResponse(address string) error {
	log.Infof("send vote response to : [%s] ", address)
	return nil
}
