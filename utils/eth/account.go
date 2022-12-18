package eth

import "eth_data_dir/utils"

func GetAccountList() (accounts []string, err error) {
	err = utils.Client.Call(&accounts, "eth_accounts")
	return
}

func GetAccountBalance(account string) (balance string, err error) {
	err = utils.Client.Call("eth_getBalance", account, "latest")
	return
}
