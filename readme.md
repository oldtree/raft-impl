## understand raft sample 

```go
package main

import (
	"eco/raft"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

var Endpoint1 = `
{
	"address":"0.0.0.0:8861",
	"peerlist":[
		{},
		{}
	],
	"name":"raft1",
	"id":1,
	"vote_timeout":500,
	"heartbeat_timeout":2000,
	"log_dir_path":"/raft/log",
	"snapshot_path":"/raft/snapshot",
	"snap_gap":4096
}`

var Endpoint2 = `
{
	"address":"0.0.0.0:8862",
	"peerlist":[
		{},
		{}
	],
	"name":"raft2",
	"id":2,
	"vote_timeout":500,
	"heartbeat_timeout":2000,
	"log_dir_path":"/raft/log",
	"snapshot_path":"/raft/snapshot",
	"snap_gap":4096
}`

var Endpoint3 = `
{
	"address":"0.0.0.0:8863",
	"peerlist":[
		{},
		{}
	],
	"name":"raft3",
	"id":3,
	"vote_timeout":500,
	"heartbeat_timeout":2000,
	"log_dir_path":"/raft/log",
	"snapshot_path":"/raft/snapshot",
	"snap_gap":4096
}`

func main() {
	c1, err := raft.NewConfigFromByte([]byte(Endpoint1))
	c2, err := raft.NewConfigFromByte([]byte(Endpoint2))
	c3, err := raft.NewConfigFromByte([]byte(Endpoint3))
	if err != nil {
		log.Error("load config failed : ", err.Error())
		os.Exit(-1)
	}
	raft1 := raft.NewRaft(c1)
	raft1.Init()
	raft2 := raft.NewRaft(c2)
	raft2.Init()
	raft3 := raft.NewRaft(c3)
	raft3.Init()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	<-sc
}

```