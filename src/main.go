package main

import (
	"MGL/api"
	"MGL/global"
	"MGL/service"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Init Logger
	api.Load_logger(0)
	// Create an instance of the app structure
	App := service.NewApp()
	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Muilt Game Launcher",
		Width:  1080,
		Height: 610,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        App.Startup,
		Frameless:        true,
		Bind: []interface{}{
			App,
		},
	})

	if err != nil {
		global.Log.Errorf("Error: %v", err.Error())
	}
}
