package main

import (
	"github.com/Anveena/RoomOfRequirement/ezConfig"
	"github.com/Anveena/RoomOfRequirement/ezLog"
	"github.com/lezhigb/KanBanMSG/httpServer"
	"github.com/lezhigb/KanBanMSG/wekanConf"
)

func main() {
	if rs, err := wekanConf.MakePasswordBase64Str("sxb_33224"); err == nil {
		if rs == "" {

		}
	}
	if err := ezConfig.ReadConf(wekanConf.Config()); err != nil {
		println(err.Error())
		return
	}
	if err := wekanConf.Config().Check(); err != nil {
		println(err.Error())
		return
	}
	if err := ezLog.SetUpEnv(&wekanConf.Config().LogConf); err != nil {
		println(err.Error())
		return
	}
	if err := httpServer.StartHTTPServer(); err != nil {
		println(err.Error())
		return
	}
}
