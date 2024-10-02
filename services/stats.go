package services

import (
	"gwentgg/db/models"
	"gwentgg/enums"
)

type RankedStats struct {
	Continent           string               `json:"continent"`
	ContinentalPosition int                  `json:"continental_position"`
	Country             string               `json:"country"`
	DateCreated         string               `json:"date_created"`
	DateLevelUpdated    string               `json:"date_level_updated"`
	DateScoreUpdated    string               `json:"date_score_updated"`
	DrawsCount          int                  `json:"draws_count"`
	FactionGamesStats   []FactionGameStats   `json:"faction_games_stats"`
	FactionProgressions []FactionProgression `json:"faction_progressions"`
	GamesCount          int                  `json:"games_count"`
	ID                  string               `json:"id"`
	Leaderboard         string               `json:"leaderboard"`
	Level               int                  `json:"level"`
	LossesCount         int                  `json:"losses_count"`
	Paragon             Paragon              `json:"paragon"`
	Platform            string               `json:"platform"`
	Position            int                  `json:"position"`
	RankID              int                  `json:"rank_id"`
	RankProgression     RankProgression      `json:"rank_progression"`
	Rating              Rating               `json:"rating"`
	Requirements        []interface{}        `json:"requirements"`
	Score               int                  `json:"score"`
	Title               string               `json:"title"`
	Username            string               `json:"username"`
	Vanities            []Vanity             `json:"vanities"`
	WinsCount           int                  `json:"wins_count"`
}

type FactionGameStats struct {
	Faction      string       `json:"faction"`
	FactionStats FactionStats `json:"faction_games_stats"`
}

func (stats FactionGameStats) ToModel(userID string) (db.FactionGameStats, error) {
	model := db.FactionGameStats{
		UserID:      userID,
		Faction:     enums.Faction(stats.Faction),
		WinsCount:   uint(stats.FactionStats.WinsCount),
		LossesCount: uint(stats.FactionStats.LossesCount),
		DrawsCount:  uint(stats.FactionStats.DrawsCount),
		GamesCount:  uint(stats.FactionStats.GamesCount),
	}
	return model, nil
}

type FactionStats struct {
	DrawsCount  int `json:"draws_count"`
	GamesCount  int `json:"games_count"`
	LossesCount int `json:"losses_count"`
	WinsCount   int `json:"wins_count"`
}

type FactionProgression struct {
	Faction            string             `json:"faction"`
	FactionProgDetails FactionProgDetails `json:"faction_progression"`
}

func (prog FactionProgression) ToModel(userID string) (db.FactionProgression, error) {
	model := db.FactionProgression{
		UserID:                           userID,
		Faction:                          enums.Faction(prog.Faction),
		IsUsedForScoreCalculation:        prog.FactionProgDetails.IsUsedForScoreCalculation,
		UnlockedScore:                    prog.FactionProgDetails.UnlockedScore,
		RealScore:                        prog.FactionProgDetails.RealScore,
		RealScorePeak:                    prog.FactionProgDetails.RealScorePeak,
		GamesCount:                       uint(prog.FactionProgDetails.GamesCount),
		UnlockedScoreGamesCountThreshold: uint(prog.FactionProgDetails.UnlockedScoreGamesCountThreshold),
		UnlockedScoreFraction:            prog.FactionProgDetails.UnlockedScoreFraction,
	}
	return model, nil
}

type FactionProgDetails struct {
	GamesCount                       int     `json:"games_count"`
	IsUsedForScoreCalculation        bool    `json:"is_used_for_score_calculation"`
	RealScore                        int     `json:"real_score"`
	RealScorePeak                    int     `json:"real_score_peak"`
	UnlockedScore                    int     `json:"unlocked_score"`
	UnlockedScoreFraction            float64 `json:"unlocked_score_fraction"`
	UnlockedScoreGamesCountThreshold int     `json:"unlocked_score_games_count_threshold"`
}

type Paragon struct {
	ParagonLevel int `json:"paragon_level"`
	PlayerLevel  int `json:"player_level"`
}

type RankProgression struct {
	MosaicPieceCount int `json:"mosaic_piece_count"`
	Rank             int `json:"rank"`
}

type Rating struct {
	Score int `json:"score"`
}

type Vanity struct {
	Category         string `json:"category"`
	ItemDefinitionID string `json:"item_definition_id"`
}

func (stats RankedStats) ToModel() (db.User, error) {
	var progressions []db.FactionProgression
	for _, prog := range stats.FactionProgressions {
		model, _ := prog.ToModel(stats.ID)
		progressions = append(progressions, model)
	}

	var factionStats []db.FactionGameStats
	for _, fgs := range stats.FactionGamesStats {
		model, _ := fgs.ToModel(stats.ID)
		factionStats = append(factionStats, model)
	}

	created, err := ParseDate(stats.DateCreated)
	if err != nil {
		return db.User{}, err
	}

	scoreUpdated, err := ParseDate(stats.DateScoreUpdated)
	if err != nil {
		return db.User{}, err
	}

	levelUpdated, err := ParseDate(stats.DateLevelUpdated)
	if err != nil {
		return db.User{}, err
	}

	var avatarID, borderID, titleID string
	for _, vanity := range stats.Vanities {
		switch vanity.Category {
		case "Avatar":
			avatarID = vanity.ItemDefinitionID
		case "Border":
			borderID = vanity.ItemDefinitionID
		case "Title":
			titleID = vanity.ItemDefinitionID
		}
	}

	user := db.User{
		ID:                  stats.ID,
		Username:            stats.Username,
		Platform:            stats.Platform,
		Title:               stats.Title,
		Score:               stats.Score,
		Position:            uint(stats.Position),
		ContinentalPosition: uint(stats.ContinentalPosition),
		RankID:              uint(stats.RankID),
		Leaderboard:         stats.Leaderboard,
		DateScoreUpdated:    scoreUpdated,
		DateLevelUpdated:    levelUpdated,
		DateCreated:         created,
		WinsCount:           uint(stats.WinsCount),
		LossesCount:         uint(stats.LossesCount),
		DrawsCount:          uint(stats.DrawsCount),
		GamesCount:          uint(stats.GamesCount),
		Level:               uint(stats.Level),
		Continent:           stats.Continent,
		Country:             stats.Country,
		MosaicPieceCount:    uint(stats.RankProgression.MosaicPieceCount),
		Rank:                uint(stats.RankProgression.Rank),
		RatingScore:         stats.Rating.Score,
		ParagonLevel:        stats.Paragon.ParagonLevel,
		AvatarID:            avatarID,
		BorderID:            borderID,
		TitleID:             titleID,
		Progressions:        progressions,
		FactionStats:        factionStats,
	}
	return user, nil
}
