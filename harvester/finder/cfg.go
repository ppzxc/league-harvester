package finder

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"time"
)

const UseWmic = "finder-use-wmic"
const TimeOut = "finder-timeout"

func Flags(cmd *cobra.Command) {
	cmd.Flags().BoolP(UseWmic, "w", false, "wmic 커맨드 사용하기")
	cmd.Flags().DurationP(TimeOut, "t", 2*time.Second, "process finder 타임아웃")
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		panic(err)
	}
}

func Config() Option {
	return Option{
		UseWmic: viper.GetBool(UseWmic),
		Timeout: viper.GetDuration(TimeOut),
	}
}
