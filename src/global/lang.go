package global

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

var (
	Path_Lang     = "Lang/"
	Use_Lang_Path = "default"
	All_Lang_Path = []string{}
)

type LanguageInfo struct {
	LanguageName        string `json:"language_name"`
	LanguageCode        string `json:"language_code"`
	TextmapPath         string `json:"textmap_path"`
	TranslationProgress string `json:"translation_progress"`
	Translator          string `json:"translator"`
	LastUpdated         string `json:"last_updated"`
	Version             string `json:"version"`
}

type LanguagePack struct {
	LanguageInfo
	Textmap map[string]string `json:"textmap"`
}

func InitLang() {
	useLang := GlobalConfig.Language
	var defaultLangPath string = "default"
	found := false
	err := filepath.WalkDir(Path_Lang, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			Log.Warnf("读取语言文件夹失败: %v", err)
			return nil
		}

		if !d.IsDir() {
			return nil
		}

		if path == Path_Lang {
			return nil
		}

		infoPath := filepath.Join(path, "info.json")
		if _, err := os.Stat(infoPath); os.IsNotExist(err) {
			Log.Warnf("跳过目录 %s: 缺少 info.json", path)
			return nil
		}

		data, err := os.ReadFile(infoPath)
		if err != nil {
			Log.Warnf("无法读取 %s: %v", infoPath, err)
			return nil
		}

		var info LanguageInfo
		if err := json.Unmarshal(data, &info); err != nil {
			Log.Warnf("无法解析 %s: %v", infoPath, err)
			return nil
		}

		if info.LanguageCode == "" || info.TextmapPath == "" {
			Log.Warnf("格式不合法: %s", path)
			return nil
		}
		All_Lang_Path = append(All_Lang_Path, info.LanguageCode)
		Log.Debugf("识别到语言: %s", info.LanguageCode)
		if info.LanguageCode == useLang {
			Log.Infof("找到匹配语言: %s", info.LanguageCode)
			Use_Lang_Path = filepath.Base(path)
			found = true
			return fmt.Errorf("Over")
		}

		return nil
	})

	if err != nil && err.Error() != "Over" {
		Log.Warnf("遍历 Lang 目录出错: %v", err)
	}

	if !found {
		Log.Infof("未找到匹配的语言，使用默认语言: %s", defaultLangPath)
		Use_Lang_Path = defaultLangPath
	}
}
