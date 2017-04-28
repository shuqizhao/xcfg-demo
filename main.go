package main

import (
	"fmt"
	"time"

	"github.com/shuqizhao/xcfg"
)

func main() {

	fmt.Println(xcfg.GetAppName())
	fmt.Println(xcfg.GetEnvironment())
	fmt.Println(xcfg.GetRemoteCfgUrl())
	ch := make(chan int)
	go func() {
		for {
			time.Sleep(time.Second * 1)
			rcfg := NewRabbitMqConfig()
			//fmt.Println(rcfg.MajorVersion, rcfg.MinorVersion)
			for i, v := range rcfg.Servers {
				fmt.Println(i, v.Server, "ps1")
			}
			time.Sleep(time.Second * 25)
			wcfg := NewRecruitWechatNotice()
			//fmt.Println(wcfg.MajorVersion, wcfg.MinorVersion)
			for i, v := range wcfg.SenderConfigurations {
				fmt.Println(i, v.SenderConfiguration, "ps2")
			}
		}
	}()

	<-ch
	fmt.Println("over")

}
