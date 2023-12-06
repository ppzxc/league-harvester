package api

import (
	log "github.com/sirupsen/logrus"
	"league-havester/harvester/model/eogStatsBlock"
)

func Request(block *eogStatsBlock.EogStatsBlock) error {
	log.Info(block)
	for i := range block.Data.Teams {
		if block.Data.Teams[i].IsPlayerTeam && block.Data.Teams[i].IsWinningTeam {
			log.Info("winning.")
		} else {
			log.Info("lose.")
		}
	}
	return nil
}
