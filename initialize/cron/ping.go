package cron

import (
	"fmt"
	"time"

	"github.com/zhangshanwen/splie/initialize/conf"
)

// 打点器 定时心跳检测 地址从内存读取降低读库压力
func Ping() {
	ticker := time.NewTicker(time.Second * time.Duration(conf.C.Distributed.Setting.PingDuration))
	go func() {
		for t := range ticker.C {
			fmt.Println(t)
		}
	}()

}
