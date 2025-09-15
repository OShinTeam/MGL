package service

import (
	"MGL/api"
	"MGL/global"
	"context"
	"encoding/json"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	runtime.WindowSetMinSize(ctx, 900, 600)
}

func (a *App) Load_logger(debug int) {
	api.Load_logger(debug)
}

func (a *App) InitConfig() {
	err := global.InitConfig()
	if err != nil {
		if global.Log != nil {
			global.Log.Fatalf("加载配置文件失败: %v", err)
		} else {
			fmt.Printf("加载配置文件失败:%v", err)
		}
	}
}

func (a *App) GetLang() string {
	language_pack, err := api.GetLang()
	if err != nil {
		global.Log.Warnf("获取Textmap失败: %v", err)
		return "1"
	}
	jsonBytes, err := json.Marshal(language_pack)
	global.Log.Debugf("%v", language_pack)
	if err != nil {
		global.Log.Warnln("Json格式出错")
		return "1"
	}
	return string(jsonBytes)
}

func (a *App) GetConfig() string {
	global.Log.Debugf("%v", global.GlobalConfig)
	jsonBytes, err := json.Marshal(global.GlobalConfig)
	if err != nil {
		global.Log.Warnln("Json格式出错")
		return "1"
	}
	return string(jsonBytes)
}
