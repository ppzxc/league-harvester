package api

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Flags(cmd *cobra.Command) {
	cmd.Flags().StringP("server-protocol", "C", "http", "서버 프로토콜")
	cmd.Flags().StringP("server-host", "e", "localhost", "서버 호스트")
	cmd.Flags().IntP("server-port", "P", 8487, "서버 포트")
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		panic(err)
	}
}

func Config() Option {
	return Option{
		Protocol: viper.GetString("server-protocol"),
		Host:     viper.GetString("server-host"),
		Port:     viper.GetInt("server-port"),
	}
}
