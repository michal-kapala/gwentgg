package models

import (
	"fmt"
	"time"
	"gwentgg/enums"

	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	ID                 string `gorm:"primaryKey"`
	Type               enums.GameType
	DateStarted        time.Time
	DateFinished       time.Time
	Finished           enums.GameEnd
	Winner             *string
	GoodGameAllowed    bool
	Valid              bool
	GameVersion        string
	VotingPatchVersion string
	DataVersion        string
	LogicVersion       string
	PackageVersion     string
	Server             string
	ServerGameID       uint
	StartingPlayer     *uint
	FinishedRounds     uint
	Players            []GamePlayer `gorm:"foreignKey:GameID"`
}

func (game Game) GetPlayerFaction(playerID string) enums.Faction {
	player := game.Players[0]
	if player.PlayerID != playerID {
		player = game.Players[1]
	}
	if player.DeckFaction == nil {
		return enums.Neutral
	}
	return enums.Faction(*player.DeckFaction)
}

func (game Game) GetOpponent(playerID string) *GamePlayer {
	player := game.Players[0]
	if player.PlayerID == playerID {
		player = game.Players[1]
	}
	return &player
}

func (game Game) DidOpen(playerID string) bool {
	if game.StartingPlayer == nil {
		return false
	}
	return fmt.Sprintf("%d", *game.StartingPlayer) == playerID
}

func GetFactionStats(playerID string, faction enums.Faction, games []Game) FactionGameStats {
	stats := FactionGameStats{UserID: playerID, Faction: faction}
	for _, game := range games {
		if faction == game.GetPlayerFaction(playerID) {
			if game.Winner == nil {
				stats.DrawsCount += 1
			} else if *game.Winner == playerID {
				stats.WinsCount += 1
			} else {
				stats.LossesCount += 1
			}
			stats.GamesCount += 1
		}
	}
	return stats
}

func GetAllFactionStats(playerID string, games []Game) []FactionGameStats {
	stats := []FactionGameStats{}
	stats = append(stats, GetFactionStats(playerID, enums.Monsters, games))
	stats = append(stats, GetFactionStats(playerID, enums.Nilfgaard, games))
	stats = append(stats, GetFactionStats(playerID, enums.NorthernKingdoms, games))
	stats = append(stats, GetFactionStats(playerID, enums.Scoiatael, games))
	stats = append(stats, GetFactionStats(playerID, enums.Skellige, games))
	stats = append(stats, GetFactionStats(playerID, enums.Syndicate, games))
	return stats
}

func GetMostPlayedFaction(stats []FactionGameStats) enums.Faction {
	factionIdx := 0
	games := stats[0].GamesCount
	for idx := 1; idx < 6; idx++ {
		if stats[idx].GamesCount > games {
			games = stats[idx].GamesCount
			factionIdx = idx
		}
	}

	if games == 0 {
		return enums.Neutral
	}
	return stats[factionIdx].Faction
}

func AggregateStats(stats []FactionGameStats) *FactionGameStats {
	result := FactionGameStats{Faction: enums.Neutral}
	for _, faction := range stats {
		result.WinsCount += faction.WinsCount
		result.LossesCount += faction.LossesCount
		result.DrawsCount += faction.DrawsCount
		result.GamesCount += faction.GamesCount
	}
	return &result
}
