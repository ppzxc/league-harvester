package cmd

import (
	"league-havester/connector"
	"league-havester/finder"
	"league-havester/harvester"
	"league-havester/logging"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "실시간 실행하기",
	Run: func(cmd *cobra.Command, args []string) {
		logging.Setup(logging.Config())
		harvester.New(harvester.Config(), finder.Config()).Start()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	logging.Flags(runCmd)
	harvester.Flags(runCmd)
	finder.Flags(runCmd)
	connector.Flags(runCmd)
}
