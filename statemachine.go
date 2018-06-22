package raft

import log "github.com/sirupsen/logrus"

type StateMachine struct {
	State      RaftState
	LastTerm   int64
	LastCommit int64
}

//If followers don't hear from a leader then they can become a candidate.

//The candidate becomes the leader if it gets votes from a majority of nodes.
func (statemachine *StateMachine) TransStatus(targetState RaftState, condition interface{}, LastTerm, LastCommit int64) (RaftState, error) {
	log.Infof("TransStatus from [%d] to [%d],condition : [%v] at term : [%d] commit : [%d]", statemachine.State, targetState, condition, LastTerm, LastCommit)
	if statemachine.State == targetState {
		return statemachine.State, nil
	}
	return statemachine.State, nil
}
