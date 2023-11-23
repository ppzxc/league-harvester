package harvester

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"time"
)

func Flags(cmd *cobra.Command) {
	cmd.Flags().Duration("harvester-duration", 5*time.Second, "수집기 메인 실행 반복 시간")
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		panic(err)
	}
}

func Config() Option {
	return Option{
		HarvesterDuration: viper.GetDuration("harvester-duration"),
	}
}
