package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"league-havester/harvester/model/eogStatsBlock"
)

var client = resty.New()

func Request(block *eogStatsBlock.EogStatsBlock, option Option) error {
	log.WithField("json", prettyPrint(block)).Debug("call")
	isWinning := isWinningCheck(block)
	winLogging(isWinning)
	if block.EventType == "Create" {
		body := of(block, isWinning)
		resp, err := client.R().
			SetHeader("Content-Type", "application/json;charset=UTF8").
			SetHeader("Accept", "application/json").
			SetBody(body).
			SetResult(map[string]interface{}{}).
			SetError(map[string]interface{}{}).
			Post(fmt.Sprintf("%s://%s:%d/api/v1/players/%s/games", option.Protocol, option.Host, option.Port, body.Puuid))
		logging(resp, err)
		if err != nil {
			return errors.New(resp.String())
		} else {
			return nil
		}
	}
	return nil
}

func isWinningCheck(block *eogStatsBlock.EogStatsBlock) bool {
	isWinning := false
	for i := range block.Data.Teams {
		if block.Data.Teams[i].IsPlayerTeam {
			if block.Data.Teams[i].IsWinningTeam {
				isWinning = true
				break
			}
		}
	}
	return isWinning
}

func winLogging(isWinning bool) {
	if isWinning {
		log.WithField("isWinning", isWinning).Debug("win")
	} else {
		log.WithField("isWinning", isWinning).Debug("lose")
	}
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func of(block *eogStatsBlock.EogStatsBlock, winning bool) *EndOfGameBlockBody {
	return &EndOfGameBlockBody{
		EndOfGameBlockEventType: block.EventType,
		EndOfGameBlockUri:       block.URI,
		GameID:                  block.Data.GameID,
		ReportGameID:            block.Data.ReportGameID,
		GameMode:                block.Data.GameMode,
		GameType:                block.Data.GameType,
		DetectedTeamPosition:    block.Data.LocalPlayer.DetectedTeamPosition,
		SelectedPosition:        block.Data.LocalPlayer.SelectedPosition,
		QueueType:               block.Data.QueueType,
		Ranked:                  block.Data.Ranked,
		ChampionID:              int64(block.Data.LocalPlayer.ChampionID),
		ChampionName:            block.Data.LocalPlayer.ChampionName,
		Spell1ID:                int64(block.Data.LocalPlayer.Spell1ID),
		Spell2ID:                int64(block.Data.LocalPlayer.Spell2ID),
		Leaver:                  block.Data.LocalPlayer.Leaver,
		Stats: Stats{
			Assists:                      int64(block.Data.LocalPlayer.Stats.Assists),
			BarracksKilled:               int64(block.Data.LocalPlayer.Stats.BarracksKilled),
			ChampionsKilled:              int64(block.Data.LocalPlayer.Stats.ChampionsKilled),
			GameEndedInEarlySurrender:    int64(block.Data.LocalPlayer.Stats.GameEndedInEarlySurrender),
			GameEndedInSurrender:         int64(block.Data.LocalPlayer.Stats.GameEndedInSurrender),
			GoldEarned:                   int64(block.Data.LocalPlayer.Stats.GoldEarned),
			NumDeaths:                    int64(block.Data.LocalPlayer.Stats.NumDeaths),
			TotalDamageDealt:             int64(block.Data.LocalPlayer.Stats.TotalDamageDealt),
			TotalDamageDealtToBuildings:  int64(block.Data.LocalPlayer.Stats.TotalDamageDealtToBuildings),
			TotalDamageDealtToChampions:  int64(block.Data.LocalPlayer.Stats.MagicDamageDealtToChampions),
			TotalDamageDealtToObjectives: int64(block.Data.LocalPlayer.Stats.TotalDamageDealtToObjectives),
			TotalDamageDealtToTurrets:    int64(block.Data.LocalPlayer.Stats.TotalDamageDealtToTurrets),
			TotalDamageTaken:             int64(block.Data.LocalPlayer.Stats.TotalDamageTaken),
		},
		Puuid:        block.Data.LocalPlayer.Puuid,
		SummonerID:   int64(block.Data.LocalPlayer.SummonerID),
		SummonerName: block.Data.LocalPlayer.SummonerName,
		BotPlayer:    block.Data.LocalPlayer.BotPlayer,
		LocalPlayer:  block.Data.LocalPlayer.IsLocalPlayer,
		Winning:      winning,
	}
}

func logging(resp *resty.Response, err error) {
	log.Debug("Response Info")
	log.WithField("Error", err).Debug()
	log.WithField("Status Code", resp.StatusCode()).Debug()
	log.WithField("Status", resp.Status()).Debug()
	log.WithField("Proto", resp.Proto()).Debug()
	log.WithField("Time", resp.Time()).Debug()
	log.WithField("Received At", resp.ReceivedAt()).Debug()
	log.WithField("Body", resp).Debug()

	// Explore trace info
	log.Debug("Request Trace Info")
	ti := resp.Request.TraceInfo()
	log.WithField("DNSLookup", ti.DNSLookup).Debug()
	log.WithField("ConnTime", ti.ConnTime).Debug()
	log.WithField("TCPConnTime", ti.TCPConnTime).Debug()
	log.WithField("TLSHandshake", ti.TLSHandshake).Debug()
	log.WithField("ServerTime", ti.ServerTime).Debug()
	log.WithField("ResponseTime", ti.ResponseTime).Debug()
	log.WithField("TotalTime", ti.TotalTime).Debug()
	log.WithField("IsConnReused", ti.IsConnReused).Debug()
	log.WithField("IsConnWasIdle", ti.IsConnWasIdle).Debug()
	log.WithField("ConnIdleTime", ti.ConnIdleTime).Debug()
	log.WithField("RequestAttempt", ti.RequestAttempt).Debug()
	log.WithField("RemoteAddr", ti.RemoteAddr).Debug()
}
