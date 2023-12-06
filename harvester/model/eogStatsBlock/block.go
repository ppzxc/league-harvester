package eogStatsBlock

type EogStatsBlock struct {
	Data struct {
		BasePoints                int      `json:"basePoints"`
		BattleBoostIPEarned       int      `json:"battleBoostIpEarned"`
		BoostIPEarned             int      `json:"boostIpEarned"`
		BoostXpEarned             int      `json:"boostXpEarned"`
		CausedEarlySurrender      bool     `json:"causedEarlySurrender"`
		CurrentLevel              int      `json:"currentLevel"`
		Difficulty                string   `json:"difficulty"`
		EarlySurrenderAccomplice  bool     `json:"earlySurrenderAccomplice"`
		ExperienceEarned          int      `json:"experienceEarned"`
		ExperienceTotal           int      `json:"experienceTotal"`
		FirstWinBonus             int      `json:"firstWinBonus"`
		GameEndedInEarlySurrender bool     `json:"gameEndedInEarlySurrender"`
		GameID                    int64    `json:"gameId"`
		GameLength                int      `json:"gameLength"`
		GameMode                  string   `json:"gameMode"`
		GameMutators              []string `json:"gameMutators"`
		GameType                  string   `json:"gameType"`
		GlobalBoostXpEarned       int      `json:"globalBoostXpEarned"`
		Invalid                   bool     `json:"invalid"`
		IPEarned                  int      `json:"ipEarned"`
		IPTotal                   int      `json:"ipTotal"`
		LeveledUp                 bool     `json:"leveledUp"`
		LocalPlayer               struct {
			BotPlayer                  bool   `json:"botPlayer"`
			ChampionID                 int    `json:"championId"`
			ChampionName               string `json:"championName"`
			ChampionSquarePortraitPath string `json:"championSquarePortraitPath"`
			DetectedTeamPosition       string `json:"detectedTeamPosition"`
			GameID                     int64  `json:"gameId"`
			IsLocalPlayer              bool   `json:"isLocalPlayer"`
			Items                      []int  `json:"items"`
			Leaver                     bool   `json:"leaver"`
			Leaves                     int    `json:"leaves"`
			Level                      int    `json:"level"`
			Losses                     int    `json:"losses"`
			ProfileIconID              int    `json:"profileIconId"`
			Puuid                      string `json:"puuid"`
			SelectedPosition           string `json:"selectedPosition"`
			SkinEmblemPaths            []any  `json:"skinEmblemPaths"`
			SkinSplashPath             string `json:"skinSplashPath"`
			SkinTilePath               string `json:"skinTilePath"`
			Spell1ID                   int    `json:"spell1Id"`
			Spell2ID                   int    `json:"spell2Id"`
			Stats                      struct {
				Assists                        int `json:"ASSISTS"`
				BarracksKilled                 int `json:"BARRACKS_KILLED"`
				ChampionsKilled                int `json:"CHAMPIONS_KILLED"`
				GameEndedInEarlySurrender      int `json:"GAME_ENDED_IN_EARLY_SURRENDER"`
				GameEndedInSurrender           int `json:"GAME_ENDED_IN_SURRENDER"`
				GoldEarned                     int `json:"GOLD_EARNED"`
				LargestCriticalStrike          int `json:"LARGEST_CRITICAL_STRIKE"`
				LargestKillingSpree            int `json:"LARGEST_KILLING_SPREE"`
				LargestMultiKill               int `json:"LARGEST_MULTI_KILL"`
				Level                          int `json:"LEVEL"`
				Lose                           int `json:"LOSE"`
				MagicDamageDealtPlayer         int `json:"MAGIC_DAMAGE_DEALT_PLAYER"`
				MagicDamageDealtToChampions    int `json:"MAGIC_DAMAGE_DEALT_TO_CHAMPIONS"`
				MagicDamageTaken               int `json:"MAGIC_DAMAGE_TAKEN"`
				MinionsKilled                  int `json:"MINIONS_KILLED"`
				NeutralMinionsKilled           int `json:"NEUTRAL_MINIONS_KILLED"`
				NumDeaths                      int `json:"NUM_DEATHS"`
				Perk0                          int `json:"PERK0"`
				Perk0Var1                      int `json:"PERK0_VAR1"`
				Perk0Var2                      int `json:"PERK0_VAR2"`
				Perk0Var3                      int `json:"PERK0_VAR3"`
				Perk1                          int `json:"PERK1"`
				Perk1Var1                      int `json:"PERK1_VAR1"`
				Perk1Var2                      int `json:"PERK1_VAR2"`
				Perk1Var3                      int `json:"PERK1_VAR3"`
				Perk2                          int `json:"PERK2"`
				Perk2Var1                      int `json:"PERK2_VAR1"`
				Perk2Var2                      int `json:"PERK2_VAR2"`
				Perk2Var3                      int `json:"PERK2_VAR3"`
				Perk3                          int `json:"PERK3"`
				Perk3Var1                      int `json:"PERK3_VAR1"`
				Perk3Var2                      int `json:"PERK3_VAR2"`
				Perk3Var3                      int `json:"PERK3_VAR3"`
				Perk4                          int `json:"PERK4"`
				Perk4Var1                      int `json:"PERK4_VAR1"`
				Perk4Var2                      int `json:"PERK4_VAR2"`
				Perk4Var3                      int `json:"PERK4_VAR3"`
				Perk5                          int `json:"PERK5"`
				Perk5Var1                      int `json:"PERK5_VAR1"`
				Perk5Var2                      int `json:"PERK5_VAR2"`
				Perk5Var3                      int `json:"PERK5_VAR3"`
				PerkPrimaryStyle               int `json:"PERK_PRIMARY_STYLE"`
				PerkSubStyle                   int `json:"PERK_SUB_STYLE"`
				PhysicalDamageDealtPlayer      int `json:"PHYSICAL_DAMAGE_DEALT_PLAYER"`
				PhysicalDamageDealtToChampions int `json:"PHYSICAL_DAMAGE_DEALT_TO_CHAMPIONS"`
				PhysicalDamageTaken            int `json:"PHYSICAL_DAMAGE_TAKEN"`
				PlayerAugment1                 int `json:"PLAYER_AUGMENT_1"`
				PlayerAugment2                 int `json:"PLAYER_AUGMENT_2"`
				PlayerAugment3                 int `json:"PLAYER_AUGMENT_3"`
				PlayerAugment4                 int `json:"PLAYER_AUGMENT_4"`
				PlayerSubteam                  int `json:"PLAYER_SUBTEAM"`
				PlayerSubteamPlacement         int `json:"PLAYER_SUBTEAM_PLACEMENT"`
				SightWardsBoughtInGame         int `json:"SIGHT_WARDS_BOUGHT_IN_GAME"`
				Spell1Cast                     int `json:"SPELL1_CAST"`
				Spell2Cast                     int `json:"SPELL2_CAST"`
				TeamEarlySurrendered           int `json:"TEAM_EARLY_SURRENDERED"`
				TeamObjective                  int `json:"TEAM_OBJECTIVE"`
				TimeCcingOthers                int `json:"TIME_CCING_OTHERS"`
				TotalDamageDealt               int `json:"TOTAL_DAMAGE_DEALT"`
				TotalDamageDealtToBuildings    int `json:"TOTAL_DAMAGE_DEALT_TO_BUILDINGS"`
				TotalDamageDealtToChampions    int `json:"TOTAL_DAMAGE_DEALT_TO_CHAMPIONS"`
				TotalDamageDealtToObjectives   int `json:"TOTAL_DAMAGE_DEALT_TO_OBJECTIVES"`
				TotalDamageDealtToTurrets      int `json:"TOTAL_DAMAGE_DEALT_TO_TURRETS"`
				TotalDamageSelfMitigated       int `json:"TOTAL_DAMAGE_SELF_MITIGATED"`
				TotalDamageShieldedOnTeammates int `json:"TOTAL_DAMAGE_SHIELDED_ON_TEAMMATES"`
				TotalDamageTaken               int `json:"TOTAL_DAMAGE_TAKEN"`
				TotalHeal                      int `json:"TOTAL_HEAL"`
				TotalHealOnTeammates           int `json:"TOTAL_HEAL_ON_TEAMMATES"`
				TotalTimeCrowdControlDealt     int `json:"TOTAL_TIME_CROWD_CONTROL_DEALT"`
				TotalTimeSpentDead             int `json:"TOTAL_TIME_SPENT_DEAD"`
				TrueDamageDealtPlayer          int `json:"TRUE_DAMAGE_DEALT_PLAYER"`
				TrueDamageDealtToChampions     int `json:"TRUE_DAMAGE_DEALT_TO_CHAMPIONS"`
				TrueDamageTaken                int `json:"TRUE_DAMAGE_TAKEN"`
				TurretsKilled                  int `json:"TURRETS_KILLED"`
				VisionWardsBoughtInGame        int `json:"VISION_WARDS_BOUGHT_IN_GAME"`
				WasAfk                         int `json:"WAS_AFK"`
			} `json:"stats"`
			SummonerID   int    `json:"summonerId"`
			SummonerName string `json:"summonerName"`
			TeamID       int    `json:"teamId"`
			Wins         int    `json:"wins"`
		} `json:"localPlayer"`
		LoyaltyBoostXpEarned int `json:"loyaltyBoostXpEarned"`
		MissionsXpEarned     int `json:"missionsXpEarned"`
		MucJwtDto            struct {
			ChannelClaim string `json:"channelClaim"`
			Domain       string `json:"domain"`
			Jwt          string `json:"jwt"`
			TargetRegion string `json:"targetRegion"`
		} `json:"mucJwtDto"`
		MultiUserChatID           string `json:"multiUserChatId"`
		MultiUserChatPassword     string `json:"multiUserChatPassword"`
		MyTeamStatus              string `json:"myTeamStatus"`
		NewSpells                 []any  `json:"newSpells"`
		NextLevelXp               int    `json:"nextLevelXp"`
		PreLevelUpExperienceTotal int    `json:"preLevelUpExperienceTotal"`
		PreLevelUpNextLevelXp     int    `json:"preLevelUpNextLevelXp"`
		PreviousLevel             int    `json:"previousLevel"`
		PreviousXpTotal           int    `json:"previousXpTotal"`
		QueueType                 string `json:"queueType"`
		Ranked                    bool   `json:"ranked"`
		ReportGameID              int64  `json:"reportGameId"`
		RerollData                struct {
			PointChangeFromChampionsOwned int `json:"pointChangeFromChampionsOwned"`
			PointChangeFromGameplay       int `json:"pointChangeFromGameplay"`
			PointsUntilNextReroll         int `json:"pointsUntilNextReroll"`
			PointsUsed                    int `json:"pointsUsed"`
			PreviousPoints                int `json:"previousPoints"`
			RerollCount                   int `json:"rerollCount"`
			TotalPoints                   int `json:"totalPoints"`
		} `json:"rerollData"`
		RpEarned  int `json:"rpEarned"`
		TeamBoost struct {
			AvailableSkins       []any  `json:"availableSkins"`
			IPReward             int    `json:"ipReward"`
			IPRewardForPurchaser int    `json:"ipRewardForPurchaser"`
			Price                int    `json:"price"`
			SkinUnlockMode       string `json:"skinUnlockMode"`
			SummonerName         string `json:"summonerName"`
			Unlocked             bool   `json:"unlocked"`
		} `json:"teamBoost"`
		TeamEarlySurrendered bool `json:"teamEarlySurrendered"`
		Teams                []struct {
			FullID             string `json:"fullId"`
			IsBottomTeam       bool   `json:"isBottomTeam"`
			IsPlayerTeam       bool   `json:"isPlayerTeam"`
			IsWinningTeam      bool   `json:"isWinningTeam"`
			MemberStatusString string `json:"memberStatusString"`
			Name               string `json:"name"`
			Players            []struct {
				BotPlayer                  bool   `json:"botPlayer"`
				ChampionID                 int    `json:"championId"`
				ChampionName               string `json:"championName"`
				ChampionSquarePortraitPath string `json:"championSquarePortraitPath"`
				DetectedTeamPosition       string `json:"detectedTeamPosition"`
				GameID                     int64  `json:"gameId"`
				IsLocalPlayer              bool   `json:"isLocalPlayer"`
				Items                      []int  `json:"items"`
				Leaver                     bool   `json:"leaver"`
				Leaves                     int    `json:"leaves"`
				Level                      int    `json:"level"`
				Losses                     int    `json:"losses"`
				ProfileIconID              int    `json:"profileIconId"`
				Puuid                      string `json:"puuid"`
				SelectedPosition           string `json:"selectedPosition"`
				SkinEmblemPaths            []any  `json:"skinEmblemPaths"`
				SkinSplashPath             string `json:"skinSplashPath"`
				SkinTilePath               string `json:"skinTilePath"`
				Spell1ID                   int    `json:"spell1Id"`
				Spell2ID                   int    `json:"spell2Id"`
				Stats                      struct {
					Assists                        int `json:"ASSISTS"`
					BarracksKilled                 int `json:"BARRACKS_KILLED"`
					ChampionsKilled                int `json:"CHAMPIONS_KILLED"`
					GameEndedInEarlySurrender      int `json:"GAME_ENDED_IN_EARLY_SURRENDER"`
					GameEndedInSurrender           int `json:"GAME_ENDED_IN_SURRENDER"`
					GoldEarned                     int `json:"GOLD_EARNED"`
					LargestCriticalStrike          int `json:"LARGEST_CRITICAL_STRIKE"`
					LargestKillingSpree            int `json:"LARGEST_KILLING_SPREE"`
					LargestMultiKill               int `json:"LARGEST_MULTI_KILL"`
					Level                          int `json:"LEVEL"`
					Lose                           int `json:"LOSE"`
					MagicDamageDealtPlayer         int `json:"MAGIC_DAMAGE_DEALT_PLAYER"`
					MagicDamageDealtToChampions    int `json:"MAGIC_DAMAGE_DEALT_TO_CHAMPIONS"`
					MagicDamageTaken               int `json:"MAGIC_DAMAGE_TAKEN"`
					MinionsKilled                  int `json:"MINIONS_KILLED"`
					NeutralMinionsKilled           int `json:"NEUTRAL_MINIONS_KILLED"`
					NumDeaths                      int `json:"NUM_DEATHS"`
					Perk0                          int `json:"PERK0"`
					Perk0Var1                      int `json:"PERK0_VAR1"`
					Perk0Var2                      int `json:"PERK0_VAR2"`
					Perk0Var3                      int `json:"PERK0_VAR3"`
					Perk1                          int `json:"PERK1"`
					Perk1Var1                      int `json:"PERK1_VAR1"`
					Perk1Var2                      int `json:"PERK1_VAR2"`
					Perk1Var3                      int `json:"PERK1_VAR3"`
					Perk2                          int `json:"PERK2"`
					Perk2Var1                      int `json:"PERK2_VAR1"`
					Perk2Var2                      int `json:"PERK2_VAR2"`
					Perk2Var3                      int `json:"PERK2_VAR3"`
					Perk3                          int `json:"PERK3"`
					Perk3Var1                      int `json:"PERK3_VAR1"`
					Perk3Var2                      int `json:"PERK3_VAR2"`
					Perk3Var3                      int `json:"PERK3_VAR3"`
					Perk4                          int `json:"PERK4"`
					Perk4Var1                      int `json:"PERK4_VAR1"`
					Perk4Var2                      int `json:"PERK4_VAR2"`
					Perk4Var3                      int `json:"PERK4_VAR3"`
					Perk5                          int `json:"PERK5"`
					Perk5Var1                      int `json:"PERK5_VAR1"`
					Perk5Var2                      int `json:"PERK5_VAR2"`
					Perk5Var3                      int `json:"PERK5_VAR3"`
					PerkPrimaryStyle               int `json:"PERK_PRIMARY_STYLE"`
					PerkSubStyle                   int `json:"PERK_SUB_STYLE"`
					PhysicalDamageDealtPlayer      int `json:"PHYSICAL_DAMAGE_DEALT_PLAYER"`
					PhysicalDamageDealtToChampions int `json:"PHYSICAL_DAMAGE_DEALT_TO_CHAMPIONS"`
					PhysicalDamageTaken            int `json:"PHYSICAL_DAMAGE_TAKEN"`
					PlayerAugment1                 int `json:"PLAYER_AUGMENT_1"`
					PlayerAugment2                 int `json:"PLAYER_AUGMENT_2"`
					PlayerAugment3                 int `json:"PLAYER_AUGMENT_3"`
					PlayerAugment4                 int `json:"PLAYER_AUGMENT_4"`
					PlayerSubteam                  int `json:"PLAYER_SUBTEAM"`
					PlayerSubteamPlacement         int `json:"PLAYER_SUBTEAM_PLACEMENT"`
					SightWardsBoughtInGame         int `json:"SIGHT_WARDS_BOUGHT_IN_GAME"`
					Spell1Cast                     int `json:"SPELL1_CAST"`
					Spell2Cast                     int `json:"SPELL2_CAST"`
					TeamEarlySurrendered           int `json:"TEAM_EARLY_SURRENDERED"`
					TeamObjective                  int `json:"TEAM_OBJECTIVE"`
					TimeCcingOthers                int `json:"TIME_CCING_OTHERS"`
					TotalDamageDealt               int `json:"TOTAL_DAMAGE_DEALT"`
					TotalDamageDealtToBuildings    int `json:"TOTAL_DAMAGE_DEALT_TO_BUILDINGS"`
					TotalDamageDealtToChampions    int `json:"TOTAL_DAMAGE_DEALT_TO_CHAMPIONS"`
					TotalDamageDealtToObjectives   int `json:"TOTAL_DAMAGE_DEALT_TO_OBJECTIVES"`
					TotalDamageDealtToTurrets      int `json:"TOTAL_DAMAGE_DEALT_TO_TURRETS"`
					TotalDamageSelfMitigated       int `json:"TOTAL_DAMAGE_SELF_MITIGATED"`
					TotalDamageShieldedOnTeammates int `json:"TOTAL_DAMAGE_SHIELDED_ON_TEAMMATES"`
					TotalDamageTaken               int `json:"TOTAL_DAMAGE_TAKEN"`
					TotalHeal                      int `json:"TOTAL_HEAL"`
					TotalHealOnTeammates           int `json:"TOTAL_HEAL_ON_TEAMMATES"`
					TotalTimeCrowdControlDealt     int `json:"TOTAL_TIME_CROWD_CONTROL_DEALT"`
					TotalTimeSpentDead             int `json:"TOTAL_TIME_SPENT_DEAD"`
					TrueDamageDealtPlayer          int `json:"TRUE_DAMAGE_DEALT_PLAYER"`
					TrueDamageDealtToChampions     int `json:"TRUE_DAMAGE_DEALT_TO_CHAMPIONS"`
					TrueDamageTaken                int `json:"TRUE_DAMAGE_TAKEN"`
					TurretsKilled                  int `json:"TURRETS_KILLED"`
					VisionWardsBoughtInGame        int `json:"VISION_WARDS_BOUGHT_IN_GAME"`
					WasAfk                         int `json:"WAS_AFK"`
				} `json:"stats"`
				SummonerID   int    `json:"summonerId"`
				SummonerName string `json:"summonerName"`
				TeamID       int    `json:"teamId"`
				Wins         int    `json:"wins"`
			} `json:"players"`
			Stats struct {
				Assists                        int `json:"ASSISTS"`
				BarracksKilled                 int `json:"BARRACKS_KILLED"`
				ChampionsKilled                int `json:"CHAMPIONS_KILLED"`
				GameEndedInEarlySurrender      int `json:"GAME_ENDED_IN_EARLY_SURRENDER"`
				GameEndedInSurrender           int `json:"GAME_ENDED_IN_SURRENDER"`
				GoldEarned                     int `json:"GOLD_EARNED"`
				LargestCriticalStrike          int `json:"LARGEST_CRITICAL_STRIKE"`
				LargestKillingSpree            int `json:"LARGEST_KILLING_SPREE"`
				LargestMultiKill               int `json:"LARGEST_MULTI_KILL"`
				Level                          int `json:"LEVEL"`
				Lose                           int `json:"LOSE"`
				MagicDamageDealtPlayer         int `json:"MAGIC_DAMAGE_DEALT_PLAYER"`
				MagicDamageDealtToChampions    int `json:"MAGIC_DAMAGE_DEALT_TO_CHAMPIONS"`
				MagicDamageTaken               int `json:"MAGIC_DAMAGE_TAKEN"`
				MinionsKilled                  int `json:"MINIONS_KILLED"`
				NeutralMinionsKilled           int `json:"NEUTRAL_MINIONS_KILLED"`
				NumDeaths                      int `json:"NUM_DEATHS"`
				Perk0                          int `json:"PERK0"`
				Perk0Var1                      int `json:"PERK0_VAR1"`
				Perk0Var2                      int `json:"PERK0_VAR2"`
				Perk0Var3                      int `json:"PERK0_VAR3"`
				Perk1                          int `json:"PERK1"`
				Perk1Var1                      int `json:"PERK1_VAR1"`
				Perk1Var2                      int `json:"PERK1_VAR2"`
				Perk1Var3                      int `json:"PERK1_VAR3"`
				Perk2                          int `json:"PERK2"`
				Perk2Var1                      int `json:"PERK2_VAR1"`
				Perk2Var2                      int `json:"PERK2_VAR2"`
				Perk2Var3                      int `json:"PERK2_VAR3"`
				Perk3                          int `json:"PERK3"`
				Perk3Var1                      int `json:"PERK3_VAR1"`
				Perk3Var2                      int `json:"PERK3_VAR2"`
				Perk3Var3                      int `json:"PERK3_VAR3"`
				Perk4                          int `json:"PERK4"`
				Perk4Var1                      int `json:"PERK4_VAR1"`
				Perk4Var2                      int `json:"PERK4_VAR2"`
				Perk4Var3                      int `json:"PERK4_VAR3"`
				Perk5                          int `json:"PERK5"`
				Perk5Var1                      int `json:"PERK5_VAR1"`
				Perk5Var2                      int `json:"PERK5_VAR2"`
				Perk5Var3                      int `json:"PERK5_VAR3"`
				PerkPrimaryStyle               int `json:"PERK_PRIMARY_STYLE"`
				PerkSubStyle                   int `json:"PERK_SUB_STYLE"`
				PhysicalDamageDealtPlayer      int `json:"PHYSICAL_DAMAGE_DEALT_PLAYER"`
				PhysicalDamageDealtToChampions int `json:"PHYSICAL_DAMAGE_DEALT_TO_CHAMPIONS"`
				PhysicalDamageTaken            int `json:"PHYSICAL_DAMAGE_TAKEN"`
				PlayerAugment1                 int `json:"PLAYER_AUGMENT_1"`
				PlayerAugment2                 int `json:"PLAYER_AUGMENT_2"`
				PlayerAugment3                 int `json:"PLAYER_AUGMENT_3"`
				PlayerAugment4                 int `json:"PLAYER_AUGMENT_4"`
				PlayerSubteam                  int `json:"PLAYER_SUBTEAM"`
				PlayerSubteamPlacement         int `json:"PLAYER_SUBTEAM_PLACEMENT"`
				SightWardsBoughtInGame         int `json:"SIGHT_WARDS_BOUGHT_IN_GAME"`
				Spell1Cast                     int `json:"SPELL1_CAST"`
				Spell2Cast                     int `json:"SPELL2_CAST"`
				TeamEarlySurrendered           int `json:"TEAM_EARLY_SURRENDERED"`
				TeamObjective                  int `json:"TEAM_OBJECTIVE"`
				TimeCcingOthers                int `json:"TIME_CCING_OTHERS"`
				TotalDamageDealt               int `json:"TOTAL_DAMAGE_DEALT"`
				TotalDamageDealtToBuildings    int `json:"TOTAL_DAMAGE_DEALT_TO_BUILDINGS"`
				TotalDamageDealtToChampions    int `json:"TOTAL_DAMAGE_DEALT_TO_CHAMPIONS"`
				TotalDamageDealtToObjectives   int `json:"TOTAL_DAMAGE_DEALT_TO_OBJECTIVES"`
				TotalDamageDealtToTurrets      int `json:"TOTAL_DAMAGE_DEALT_TO_TURRETS"`
				TotalDamageSelfMitigated       int `json:"TOTAL_DAMAGE_SELF_MITIGATED"`
				TotalDamageShieldedOnTeammates int `json:"TOTAL_DAMAGE_SHIELDED_ON_TEAMMATES"`
				TotalDamageTaken               int `json:"TOTAL_DAMAGE_TAKEN"`
				TotalHeal                      int `json:"TOTAL_HEAL"`
				TotalHealOnTeammates           int `json:"TOTAL_HEAL_ON_TEAMMATES"`
				TotalTimeCrowdControlDealt     int `json:"TOTAL_TIME_CROWD_CONTROL_DEALT"`
				TotalTimeSpentDead             int `json:"TOTAL_TIME_SPENT_DEAD"`
				TrueDamageDealtPlayer          int `json:"TRUE_DAMAGE_DEALT_PLAYER"`
				TrueDamageDealtToChampions     int `json:"TRUE_DAMAGE_DEALT_TO_CHAMPIONS"`
				TrueDamageTaken                int `json:"TRUE_DAMAGE_TAKEN"`
				TurretsKilled                  int `json:"TURRETS_KILLED"`
				VisionWardsBoughtInGame        int `json:"VISION_WARDS_BOUGHT_IN_GAME"`
				WasAfk                         int `json:"WAS_AFK"`
			} `json:"stats,omitempty"`
			Tag    string `json:"tag"`
			TeamID int    `json:"teamId"`
			Stats0 struct {
				Assists                        int `json:"ASSISTS"`
				BarracksKilled                 int `json:"BARRACKS_KILLED"`
				ChampionsKilled                int `json:"CHAMPIONS_KILLED"`
				GameEndedInEarlySurrender      int `json:"GAME_ENDED_IN_EARLY_SURRENDER"`
				GameEndedInSurrender           int `json:"GAME_ENDED_IN_SURRENDER"`
				GoldEarned                     int `json:"GOLD_EARNED"`
				LargestCriticalStrike          int `json:"LARGEST_CRITICAL_STRIKE"`
				LargestKillingSpree            int `json:"LARGEST_KILLING_SPREE"`
				LargestMultiKill               int `json:"LARGEST_MULTI_KILL"`
				Level                          int `json:"LEVEL"`
				MagicDamageDealtPlayer         int `json:"MAGIC_DAMAGE_DEALT_PLAYER"`
				MagicDamageDealtToChampions    int `json:"MAGIC_DAMAGE_DEALT_TO_CHAMPIONS"`
				MagicDamageTaken               int `json:"MAGIC_DAMAGE_TAKEN"`
				MinionsKilled                  int `json:"MINIONS_KILLED"`
				NeutralMinionsKilled           int `json:"NEUTRAL_MINIONS_KILLED"`
				NumDeaths                      int `json:"NUM_DEATHS"`
				Perk0                          int `json:"PERK0"`
				Perk0Var1                      int `json:"PERK0_VAR1"`
				Perk0Var2                      int `json:"PERK0_VAR2"`
				Perk0Var3                      int `json:"PERK0_VAR3"`
				Perk1                          int `json:"PERK1"`
				Perk1Var1                      int `json:"PERK1_VAR1"`
				Perk1Var2                      int `json:"PERK1_VAR2"`
				Perk1Var3                      int `json:"PERK1_VAR3"`
				Perk2                          int `json:"PERK2"`
				Perk2Var1                      int `json:"PERK2_VAR1"`
				Perk2Var2                      int `json:"PERK2_VAR2"`
				Perk2Var3                      int `json:"PERK2_VAR3"`
				Perk3                          int `json:"PERK3"`
				Perk3Var1                      int `json:"PERK3_VAR1"`
				Perk3Var2                      int `json:"PERK3_VAR2"`
				Perk3Var3                      int `json:"PERK3_VAR3"`
				Perk4                          int `json:"PERK4"`
				Perk4Var1                      int `json:"PERK4_VAR1"`
				Perk4Var2                      int `json:"PERK4_VAR2"`
				Perk4Var3                      int `json:"PERK4_VAR3"`
				Perk5                          int `json:"PERK5"`
				Perk5Var1                      int `json:"PERK5_VAR1"`
				Perk5Var2                      int `json:"PERK5_VAR2"`
				Perk5Var3                      int `json:"PERK5_VAR3"`
				PerkPrimaryStyle               int `json:"PERK_PRIMARY_STYLE"`
				PerkSubStyle                   int `json:"PERK_SUB_STYLE"`
				PhysicalDamageDealtPlayer      int `json:"PHYSICAL_DAMAGE_DEALT_PLAYER"`
				PhysicalDamageDealtToChampions int `json:"PHYSICAL_DAMAGE_DEALT_TO_CHAMPIONS"`
				PhysicalDamageTaken            int `json:"PHYSICAL_DAMAGE_TAKEN"`
				PlayerAugment1                 int `json:"PLAYER_AUGMENT_1"`
				PlayerAugment2                 int `json:"PLAYER_AUGMENT_2"`
				PlayerAugment3                 int `json:"PLAYER_AUGMENT_3"`
				PlayerAugment4                 int `json:"PLAYER_AUGMENT_4"`
				PlayerSubteam                  int `json:"PLAYER_SUBTEAM"`
				PlayerSubteamPlacement         int `json:"PLAYER_SUBTEAM_PLACEMENT"`
				SightWardsBoughtInGame         int `json:"SIGHT_WARDS_BOUGHT_IN_GAME"`
				Spell1Cast                     int `json:"SPELL1_CAST"`
				Spell2Cast                     int `json:"SPELL2_CAST"`
				TeamEarlySurrendered           int `json:"TEAM_EARLY_SURRENDERED"`
				TeamObjective                  int `json:"TEAM_OBJECTIVE"`
				TimeCcingOthers                int `json:"TIME_CCING_OTHERS"`
				TotalDamageDealt               int `json:"TOTAL_DAMAGE_DEALT"`
				TotalDamageDealtToBuildings    int `json:"TOTAL_DAMAGE_DEALT_TO_BUILDINGS"`
				TotalDamageDealtToChampions    int `json:"TOTAL_DAMAGE_DEALT_TO_CHAMPIONS"`
				TotalDamageDealtToObjectives   int `json:"TOTAL_DAMAGE_DEALT_TO_OBJECTIVES"`
				TotalDamageDealtToTurrets      int `json:"TOTAL_DAMAGE_DEALT_TO_TURRETS"`
				TotalDamageSelfMitigated       int `json:"TOTAL_DAMAGE_SELF_MITIGATED"`
				TotalDamageShieldedOnTeammates int `json:"TOTAL_DAMAGE_SHIELDED_ON_TEAMMATES"`
				TotalDamageTaken               int `json:"TOTAL_DAMAGE_TAKEN"`
				TotalHeal                      int `json:"TOTAL_HEAL"`
				TotalHealOnTeammates           int `json:"TOTAL_HEAL_ON_TEAMMATES"`
				TotalTimeCrowdControlDealt     int `json:"TOTAL_TIME_CROWD_CONTROL_DEALT"`
				TotalTimeSpentDead             int `json:"TOTAL_TIME_SPENT_DEAD"`
				TrueDamageDealtPlayer          int `json:"TRUE_DAMAGE_DEALT_PLAYER"`
				TrueDamageDealtToChampions     int `json:"TRUE_DAMAGE_DEALT_TO_CHAMPIONS"`
				TrueDamageTaken                int `json:"TRUE_DAMAGE_TAKEN"`
				TurretsKilled                  int `json:"TURRETS_KILLED"`
				VisionWardsBoughtInGame        int `json:"VISION_WARDS_BOUGHT_IN_GAME"`
				WasAfk                         int `json:"WAS_AFK"`
				Win                            int `json:"WIN"`
			} `json:"stats,omitempty"`
		} `json:"teams"`
		TimeUntilNextFirstWinBonus int `json:"timeUntilNextFirstWinBonus"`
		XbgpBoostXpEarned          int `json:"xbgpBoostXpEarned"`
	} `json:"data"`
	EventType string `json:"eventType"`
	URI       string `json:"uri"`
}