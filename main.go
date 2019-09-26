package main

import (
	"github.com/myadamtest/logkit"
)

func main() {
	logkit.Init("mylog", logkit.LevelDebug)

	logkit.Infof("create new test")
}
