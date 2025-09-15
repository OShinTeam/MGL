package api

import (
	"MGL/global"
	"MGL/logger"
)

func Load_logger(debug int) {
	if global.Log == nil {
		global.Log = logger.InitLogger()
		global.Log.Infoln("日志系统初始化成功")
	} else {
		if debug == 1 {
			if global.LogFile != nil {
				global.LogFile.Sync()
				global.LogFile.Close()
				global.LogFile = nil
			}
			global.Log = logger.InitDebugLogger()
			global.Log.Infoln("开启Debug模式日志成功")
			global.Log.Debugln(global.GlobalConfig)
		}
	}
}
