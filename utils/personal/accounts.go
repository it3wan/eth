package personal

import "eth_data_dir/utils"

func GetAccountList() (accounts []string, err error) {
	err = utils.Client.Call(&accounts, "personal_listAccounts")
	return
}

func NewAccount(pwd string) (account string, err error) {
	err = utils.Client.Call(account, "personal_newAccount", pwd)
	return
}
