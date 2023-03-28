package main

import (
	"github.com/aceld/zinx/znet"
	"time"
)

func main() {
	s := znet.NewServer()

	s.StartHeartBeat(3 * time.Second)

	s.Serve()
}
