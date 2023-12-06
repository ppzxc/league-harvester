package logging

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"league-havester/constant"
)

const LogUseJsonFormatter = "log-use-json-formatter"
const LogUseStdout = "log-use-stdout"
const LogLevel = "log-level"
const LogFileName = "log-filename"
const LogMaxSize = "log-max-size"
const LogMaxBackup = "log-max-backup"
const LogMaxAge = "log-max-age"
const LogCompress = "log-compress"

func Flags(cmd *cobra.Command) {
	cmd.Flags().BoolP(LogUseJsonFormatter, "f", false, "로그 포맷을 json 으로 변경하기")
	cmd.Flags().BoolP(LogUseStdout, "s", false, "로그를 실시간으로 확인, 미사용시 파일로 변경됨")
	cmd.Flags().StringP(LogLevel, "l", "info", "로그 레벨")
	cmd.Flags().StringP(LogFileName, "F", constant.LoggingFilePath("harvester.log"), "로그 파일 경로 및 이름")
	cmd.Flags().IntP(LogMaxSize, "S", 500, "로그 파일 최대 크기")
	cmd.Flags().IntP(LogMaxBackup, "b", 3, "로그 파일 최대 백업")
	cmd.Flags().IntP(LogMaxAge, "a", 7, "로그 파일 최대 보관 일수")
	cmd.Flags().BoolP(LogCompress, "c", true, "로그 파일 압축 여부")
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		panic(err)
	}
}

func Config() Option {
	return Option{
		UseJsonFormatter: viper.GetBool(LogUseJsonFormatter),
		UseStdout:        viper.GetBool(LogUseStdout),
		Level:            LevelParser(viper.GetString(LogLevel)),
		LogFilename:      viper.GetString(LogFileName),
		LogMaxSize:       viper.GetInt(LogMaxSize),
		LogMaxBackup:     viper.GetInt(LogMaxBackup),
		LogMaxAge:        viper.GetInt(LogMaxAge),
		LogCompress:      viper.GetBool(LogCompress),
	}
}
