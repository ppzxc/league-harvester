package finder

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"time"
)

func Flags(cmd *cobra.Command) {
	cmd.Flags().Bool("finder-use-wmic", false, "wmic 커맨드 사용하기")
	cmd.Flags().Duration("finder-timeout", 2*time.Second, "process finder 타임아웃")
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		panic(err)
	}
}

func Config() Option {
	return Option{
		UseWmic: viper.GetBool("finder-use-wmic"),
		Timeout: viper.GetDuration("finder-timeout"),
	}
}
