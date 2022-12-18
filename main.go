package main

import (
	_ "eth_data_dir/utils"
	"eth_data_dir/utils/eth"
	"eth_data_dir/utils/net"
	"eth_data_dir/utils/personal"
	"fmt"
)

func main() {

	networkid, err := net.GetNetWorkId()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(networkid)

	listening, err := net.GetNetIsListening()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(listening)

	peerCount, err := net.GetNetPeerCount()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(peerCount)

	accounts, err := eth.GetAccountList()
	fmt.Println(err)
	fmt.Println(accounts)

	accounts, err = personal.GetAccountList()
	fmt.Println(accounts)

	account, err := personal.NewAccount("1234556")
	fmt.Println(account)

}
