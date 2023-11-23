package logging

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Flags(cmd *cobra.Command) {
	cmd.Flags().Bool("use-json-formatter", false, "로그 포맷을 json 으로 변경하기")
	cmd.Flags().Bool("use-stdout", true, "로그를 실시간으로 확인, 미사용시 파일로 변경됨")
	cmd.Flags().String("level", "info", "로그 레벨")
	cmd.Flags().String("filename", "./harvester.log", "로그 파일 경로 및 이름")
	cmd.Flags().Int("max-size", 500, "로그 파일 최대 크기")
	cmd.Flags().Int("max-backup", 3, "로그 파일 최대 백업")
	cmd.Flags().Int("max-age", 7, "로그 파일 최대 보관 일수")
	cmd.Flags().Bool("compress", true, "로그 파일 압축 여부")
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		panic(err)
	}
}

func Config() Option {
	return Option{
		UseJsonFormatter: viper.GetBool("use-json-formatter"),
		UseStdout:        viper.GetBool("use-stdout"),
		Level:            LevelParser(viper.GetString("level")),
		LogFilename:      viper.GetString("filename"),
		LogMaxSize:       viper.GetInt("max-size"),
		LogMaxBackup:     viper.GetInt("max-backup"),
		LogMaxAge:        viper.GetInt("max-age"),
		LogCompress:      viper.GetBool("compress"),
	}
}
