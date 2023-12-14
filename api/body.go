package api

import "encoding/json"

type EndOfGameBlockBody struct {
	EndOfGameBlockEventType string `json:"endOfGameBlockEventType"`
	EndOfGameBlockUri       string `json:"endOfGameBlockUri"`
	GameID                  int64  `json:"gameId"`
	ReportGameID            int64  `json:"reportGameId"`
	GameMode                string `json:"gameMode"`
	GameType                string `json:"gameType"`
	DetectedTeamPosition    string `json:"detectedTeamPosition"`
	SelectedPosition        string `json:"selectedPosition"`
	QueueType               string `json:"queueType"`
	Ranked                  bool   `json:"ranked"`
	ChampionID              int64  `json:"championId"`
	ChampionName            string `json:"championName"`
	Spell1ID                int64  `json:"spell1Id"`
	Spell2ID                int64  `json:"spell2Id"`
	Leaver                  bool   `json:"leaver"`
	Stats                   Stats  `json:"stats"`
	Puuid                   string `json:"puuid"`
	SummonerID              int64  `json:"summonerId"`
	SummonerName            string `json:"summonerName"`
	BotPlayer               bool   `json:"botPlayer"`
	LocalPlayer             bool   `json:"localPlayer"`
	Winning                 bool   `json:"winning"`
}

type Stats struct {
	Assists                      int64 `json:"assists"`
	BarracksKilled               int64 `json:"barracksKilled"`
	ChampionsKilled              int64 `json:"championsKilled"`
	GameEndedInEarlySurrender    int64 `json:"gameEndedInEarlySurrender"`
	GameEndedInSurrender         int64 `json:"gameEndedInSurrender"`
	GoldEarned                   int64 `json:"goldEarned"`
	NumDeaths                    int64 `json:"numDeaths"`
	TotalDamageDealt             int64 `json:"totalDamageDealt"`
	TotalDamageDealtToBuildings  int64 `json:"totalDamageDealtToBuildings"`
	TotalDamageDealtToChampions  int64 `json:"totalDamageDealtToChampions"`
	TotalDamageDealtToObjectives int64 `json:"totalDamageDealtToObjectives"`
	TotalDamageDealtToTurrets    int64 `json:"totalDamageDealtToTurrets"`
	TotalDamageTaken             int64 `json:"totalDamageTaken"`
}

type Problem struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
	Unknown  map[string]*json.RawMessage
}

type Option struct {
	Protocol string
	Host     string
	Port     int
}
