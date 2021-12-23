package main

import (
	"os"
	"regexp"
)

func BuildRegex() []*regexp.Regexp {
	var regex = make([]*regexp.Regexp, 3)

	btc_regex, err := regexp.Compile(`\b(bc1|[13])[a-zA-HJ-NP-Z0-9]{26,35}\b`)
	if err != nil {
		os.Exit(0)
	}
	eth_regex, err := regexp.Compile(`"\b0x[a-fA-F0-9]{40}\b"`)
	if err != nil {
		os.Exit(0)
	}
	xmr_regex, err := regexp.Compile(`\b4([0-9]|[A-B])(.){93}\b`)
	if err != nil {
		os.Exit(0)
	}

	regex[0] = btc_regex
	regex[1] = eth_regex
	regex[2] = xmr_regex

	return regex
}
