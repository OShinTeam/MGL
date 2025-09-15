package api

import (
	"MGL/global"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func tryLoad(langPath string) *global.LanguagePack {
	infoPath := filepath.Join(langPath, "info.json")
	infoData, err := os.ReadFile(infoPath)
	if err != nil {
		return nil
	}

	var langInfo global.LanguageInfo
	if err := json.Unmarshal(infoData, &langInfo); err != nil {
		return nil
	}

	textmapPath := filepath.Join(langPath, "textmap.json")
	textmapData, err := os.ReadFile(textmapPath)
	if err != nil {
		return &global.LanguagePack{
			LanguageInfo: langInfo,
			Textmap:      make(map[string]string),
		}
	}

	var textmap map[string]string
	if err := json.Unmarshal(textmapData, &textmap); err != nil {
		return &global.LanguagePack{
			LanguageInfo: langInfo,
			Textmap:      make(map[string]string),
		}
	}

	return &global.LanguagePack{
		LanguageInfo: langInfo,
		Textmap:      textmap,
	}
}

func GetLang() (*global.LanguagePack, error) {
	langPath := filepath.Join(global.Path_Lang, global.Use_Lang_Path)
	global.Log.Debugf("使用语言包：%s", langPath)
	result := tryLoad(langPath)

	if result == nil {
		global.Log.Warnln("使用指定语言路径加载失败，尝试回退到 default 路径")
		defaultPath := filepath.Join(global.Path_Lang, "default")
		global.Log.Debugf("使用语言包：%s", defaultPath)
		result = tryLoad(defaultPath)
	}

	if result == nil {
		err := fmt.Errorf("无法加载任何语言包，请检查语言包是否存在且格式正确")
		global.Log.Fatalf("错误: %v", err)
		return nil, err
	}

	if global.Use_Lang_Path != "default" && result.LanguageCode != global.GlobalConfig.Language {
		global.Log.Warnf("警告: 实际加载的语言与配置不符（期望: %s，实际: %s）",
			global.GlobalConfig.Language, result.LanguageCode)
	}

	return result, nil
}

func GetALLLang() []string {
	return global.All_Lang_Path
}
