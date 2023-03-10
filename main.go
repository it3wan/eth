package main

import (
	"eth_data_dir/contranct"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
)

const key = `{"address":"3aa6b0ebf7453d0b19a6d264fe262721b53518e8","crypto":{"cipher":"aes-128-ctr","ciphertext":"be6400ad1e06e19031d46503a63f4b92217afee00b2b9a337e47eed39d494900","cipherparams":{"iv":"c310dd0f827ff0fcc068fd99a8d54df6"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":4096,"p":6,"r":8,"salt":"bd045bbf44070ffcf38f31e6b5fc968370898b2e16d96c1da9590826385c76e5"},"mac":"35f8fd5b75da6e6cd24f0657cadf0edf75306ceff00b9e7dfbfa1e4e70fb1843"},"id":"f6f11869-1f62-4e14-becb-7bdaa59a2d3a","version":3}`
const contractAddress = "0x4dBE43FEF09F629D3bAC5F3a77a504c577944b20"

func main() {
	conn, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// 创建storage实例
	storage, err := contranct.NewStorage(common.HexToAddress(contractAddress), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate Storage contract: %v", err)
	}

	// console中执行eth.chainId()得到chainID
	tx, err := bind.NewTransactorWithChainID(strings.NewReader(key), "", big.NewInt(0x539))
	if err != nil {
		log.Fatalf("Failed to create transaction")
	}

	// 调用storage.store()
	_, err = storage.Store(tx, big.NewInt(2))
	if err != nil {
		log.Fatalf("Failed to store number: %v", err)
	}

	// 调用storage.retrieve()
	number, err := storage.Retrieve(&bind.CallOpts{Pending: true})
	if err != nil {
		log.Fatalf("Failed to retrieve number: %v", err)
	}
	fmt.Println(number)
}
