package main

import (
	"io/ioutil"
	"strings"
)

func BuildWallet() map[string]string {
	wallet := make(map[string]string)

	configs, err := ioutil.ReadFile("configs.txt")
	if err != nil {
		wallet["btc_addr"] = "bitcoin_wallet"
		wallet["eth_addr"] = "ethereum_wallet"
		wallet["xmr_addr"] = "monero_wallet"
	} else {
		for i, line := range strings.Split(string(configs), "\n") {
			s := strings.Split(string(line), "=")
			if i >= 3 {
				break
			}
			wallet[s[0]] = s[1]
		}
	}

	return wallet
}
