package connector

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"league-havester/finder"
	"time"
)

func Flags(cmd *cobra.Command) {
	cmd.Flags().Duration("connector-timeout", 1*time.Second, "connector connect 타임아웃")
	cmd.Flags().String("connector-host", "127.0.0.1", "connector 호스트")
	cmd.Flags().Int("connector-port", 0, "connector 포트")
	cmd.Flags().Int("connector-pid", 0, "connector pid")
	cmd.Flags().String("connector-token", "", "connector 토큰")
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		panic(err)
	}
}

func Config(result *finder.Result) Option {
	port := viper.GetInt("connector-port")
	if port == 0 {
		port = result.Port
	}
	pid := viper.GetInt("connector-pid")
	if pid == 0 {
		pid = result.Pid
	}
	token := viper.GetString("connector-token")
	if len(token) == 0 {
		token = result.Token
	}
	return Option{
		ConnectTimeout: viper.GetDuration("connector-timeout"),
		Host:           viper.GetString("connector-host"),
		Port:           port,
		Pid:            pid,
		Token:          token,
	}
}
