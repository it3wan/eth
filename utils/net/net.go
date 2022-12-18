package net

import "eth_data_dir/utils"

func GetNetWorkId() (networkid string, err error) {
	err = utils.Client.Call(&networkid, "net_version")
	if err != nil {
		return networkid, err
	}
	return networkid, nil
}

func GetNetIsListening() (listening bool, err error) {
	err = utils.Client.Call(&listening, "net_listening")
	if err != nil {
		return listening, err
	}
	return listening, nil
}

func GetNetPeerCount() (count string, err error) {
	err = utils.Client.Call(&count, "net_peerCount")
	if err != nil {
		return count, err
	}
	return count, nil
}
