package logger

import (
	"testing"
	"time"
)

func Test_Console(t *testing.T) {
	// 初始化
	log := NewLog(1000)
	// 设置输出引擎
	log.SetEngine("file", `{"spilt":"size", "filename":"logs/test.log"}`)

	log.DelEngine("console")

	// 设置是否输出行号
	log.SetFuncCall(true)
/*
	log.Trac("Trac")
	log.Info("Info")
	log.Warn("Warning")
	log.Error("Error")
	log.Fatal("Fatal")*/

	// 设置log级别
	//log.SetLevel(Warning)

	timer1 := time.NewTicker(1 * time.Millisecond)
	for {
		select {
		case <-timer1.C:
			log.Trac("Trac")
			log.Info("Info")
			log.Warn("Warning")
			log.Error("Error")
			log.Fatal("Fatal")
		}
	}
}
