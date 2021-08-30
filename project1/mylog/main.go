package main

import (
	"time"

	fileLogger "github.com/project1/logger/file"
)

func main() {
	log := fileLogger.NewFileLog("./", "ceshi.log", "debug", 10*1024*1024, true)

	for {
		log.Debug("debug")
		log.Info("info")
		log.Waring("waring")
		log.Error("error")
		log.Fatal("fatal")
		time.Sleep(time.Second * 2)

	}
	log.CloseFile()
}
