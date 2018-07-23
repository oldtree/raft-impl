package raft

type LogEntry struct {
	Commited int64  `json:"commited,omitempty"`
	Terms    int64  `json:"terms,omitempty"`
	NodeId   int    `json:"node_id,omitempty"`
	Seq      int64  `json:"seq,omitempty"`
	CMD      []byte `json:"cmd,omitempty"`
}
