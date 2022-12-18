package utils

import (
	rpc "github.com/ethereum/go-ethereum/rpc"
)

var Client *rpc.Client

func init() {
	Client, _ = rpc.Dial("http://127.0.0.1:8545")
	if Client == nil {
		panic("连接出错")
	}
}
