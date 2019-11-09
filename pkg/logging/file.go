package logging

import (
	"fmt"
	"time"

	"github.com/zhangtao25/mangostreet-ser-gin/pkg/setting"
)

// 获取日志文件保存路径
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.LogSavePath)
}

// 获取日志文件的保存名
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt,
	)
}
