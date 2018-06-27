package raft

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type StateMachine struct {
	State      RaftState        `json:"state,omitempty"`
	LastTerm   int64            `json:"last_term,omitempty"`
	LastCommit int64            `json:"last_commit,omitempty"`
	Peers      map[string]*Peer `json:"peers,omitempty"`
}

//If followers don't hear from a leader then they can become a candidate.

//The candidate becomes the leader if it gets votes from a majority of nodes.
func (statemachine *StateMachine) TransState(targetState RaftState, condition interface{}, LastTerm, LastCommit int64) (RaftState, error) {
	log.Infof("TransStatus from [%d] to [%d],condition : [%v] at term : [%d] commit : [%d]", statemachine.State, targetState, condition, LastTerm, LastCommit)
	if statemachine.State == targetState {
		return statemachine.State, nil
	}
	return statemachine.State, nil
}

func (statemachine *StateMachine) MakeFilepath() string {
	return fmt.Sprintf("statemachine-%s-%d-%d.json", statemachine.State, statemachine.LastTerm, statemachine.LastCommit)
}

func (statemachine *StateMachine) SaveState() ([]byte, error) {
	stateData, err := json.Marshal(statemachine)
	if err != nil {
		log.Warnf("save state machine failed : [%s] ", err.Error())
		return nil, err
	}
	return stateData, nil
}

func (stateMachine *StateMachine) RecoverState(stateData []byte) error {
	err := json.Unmarshal(stateData, stateMachine)
	if err != nil {
		return err
	}
	return nil
}
