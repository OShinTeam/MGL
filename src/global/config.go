package global

import (
	"encoding/json"
	"os"
)

var (
	Path_Config = "Local/Config/config.json"
)

type GConfig struct {
	Language string `json:"language"`
}

type GConfigPack struct {
	Version string
	GConfig
}

var GlobalConfig = &GConfigPack{}

func InitConfig() error {
	file, err := os.ReadFile(Path_Config)
	if err != nil {
		return err
	}
	GlobalConfig = &GConfigPack{}
	err = json.Unmarshal(file, GlobalConfig)
	if err != nil {
		return err
	}
	GlobalConfig.Version = "1.0.0"
	InitLang()
	return nil
}
