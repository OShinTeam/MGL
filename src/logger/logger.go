package logger

import (
	"MGL/global"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

func InitLogger() *logrus.Logger {
	clear_log(global.GlobalRuntime_log)
	startTime := time.Now().Format("2006-01-02T15_04_05")
	fileName := global.GlobalRuntime_log.Path + global.GlobalRuntime_log.Prefix + "-" + startTime + ".log"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	global.LogFile = file
	if err != nil {
		fmt.Printf("无法打开日志文件: %v\n", err)
		os.Exit(1)
	}
	mLog := logrus.New()
	mLog.SetOutput(file)
	mLog.SetReportCaller(true)
	mLog.SetFormatter(&global.LogFormatter{})
	level := logrus.InfoLevel
	mLog.SetLevel(level)
	return mLog
}

func InitDebugLogger() *logrus.Logger {
	clear_log(global.GlobalDebug_log)
	startTime := time.Now().Format("2006-01-02T15_04_05")
	fileName := global.GlobalDebug_log.Path + global.GlobalDebug_log.Prefix + "-" + startTime + ".log"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	global.LogFile = file
	if err != nil {
		fmt.Printf("无法打开日志文件: %v\n", err)
		os.Exit(1)
	}
	mLog := logrus.New()
	mLog.SetOutput(file)
	mLog.SetReportCaller(true)
	mLog.SetFormatter(&global.LogFormatter{})
	level := logrus.DebugLevel
	mLog.SetLevel(level)
	return mLog
}

func clear_log(slog *global.GLog) {
	files, err := filepath.Glob(filepath.Join(slog.Path, "*.log"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "无法查找日志文件: %v\n", err)
		return
	}
	if len(files) <= slog.MaxLogFiles {
		fmt.Println("日志文件数量未超过上限，无需清理")
		return
	}
	type fileWithTime struct {
		filename string
		t        time.Time
	}
	var logFiles []fileWithTime

	for _, file := range files {
		base := filepath.Base(file)
		ext := strings.TrimSuffix(base, ".log")
		t, err := time.Parse("runtime-2006-01-02T15_04_05", ext)
		if err != nil {
			t, err = time.Parse("debug-2006-01-02T15_04_05", ext)
			if err != nil {
				fmt.Fprintf(os.Stderr, "跳过无法解析时间的文件: %s\n", file)
				continue
			}
		}
		logFiles = append(logFiles, fileWithTime{filename: file, t: t})
	}

	sort.Slice(logFiles, func(i, j int) bool {
		return logFiles[i].t.Before(logFiles[j].t)
	})

	deleteCount := len(logFiles) - slog.MaxLogFiles
	fmt.Printf("超过最大日志数，准备删除 %d 个最旧的日志文件:\n", deleteCount)

	for i := 0; i < deleteCount; i++ {
		file := logFiles[i].filename
		fmt.Println("删除:", file)
		if err := os.Remove(file); err != nil {
			fmt.Fprintf(os.Stderr, "删除失败: %v\n", err)
		}
	}
}
