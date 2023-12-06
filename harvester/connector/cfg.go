package connector

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"league-havester/harvester/finder"
	"time"
)

func Flags(cmd *cobra.Command) {
	cmd.Flags().DurationP("connector-timeout", "T", 1*time.Second, "connector connect 타임아웃")
	cmd.Flags().StringP("connector-host", "H", "127.0.0.1", "connector 호스트")
	cmd.Flags().IntP("connector-port", "p", 0, "connector 포트")
	cmd.Flags().IntP("connector-pid", "d", 0, "connector pid")
	cmd.Flags().StringP("connector-token", "k", "", "connector 토큰")
	cmd.Flags().DurationP("connector-write-timeout", "o", 10*time.Second, "write timeout")
	cmd.Flags().DurationP("connector-pong-timeout", "O", 60*time.Second, "pong timeout")
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
		WriteTimeout:   viper.GetDuration("connector-write-timeout"),
		PongWait:       viper.GetDuration("connector-pong-timeout"),
		PingPeriod:     (viper.GetDuration("connector-pong-timeout") * 9) / 10,
		Host:           viper.GetString("connector-host"),
		Port:           port,
		Pid:            pid,
		Token:          token,
	}
}
