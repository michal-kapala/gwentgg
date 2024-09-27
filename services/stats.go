package services

type RankedStats struct {
	Continent           string                `json:"continent"`
	ContinentalPosition int                   `json:"continental_position"`
	Country             string                `json:"country"`
	DateCreated         string                `json:"date_created"`
	DateLevelUpdated    string                `json:"date_level_updated"`
	DateScoreUpdated    string                `json:"date_score_updated"`
	DrawsCount          int                   `json:"draws_count"`
	FactionGamesStats   []FactionGameStats    `json:"faction_games_stats"`
	FactionProgressions []FactionProgression  `json:"faction_progressions"`
	GamesCount          int                   `json:"games_count"`
	ID                  string                `json:"id"`
	Leaderboard         string                `json:"leaderboard"`
	Level               int                   `json:"level"`
	LossesCount         int                   `json:"losses_count"`
	Paragon             Paragon               `json:"paragon"`
	Platform            string                `json:"platform"`
	Position            int                   `json:"position"`
	RankID              int                   `json:"rank_id"`
	RankProgression     RankProgression       `json:"rank_progression"`
	Rating              Rating                `json:"rating"`
	Requirements        []interface{}         `json:"requirements"`
	Score               int                   `json:"score"`
	Title               string                `json:"title"`
	Username            string                `json:"username"`
	Vanities            []interface{}         `json:"vanities"`
	WinsCount           int                   `json:"wins_count"`
}

type FactionGameStats struct {
	Faction          string            `json:"faction"`
	FactionStats     FactionStats      `json:"faction_games_stats"`
}

type FactionStats struct {
	DrawsCount   int `json:"draws_count"`
	GamesCount   int `json:"games_count"`
	LossesCount  int `json:"losses_count"`
	WinsCount    int `json:"wins_count"`
}

type FactionProgression struct {
	Faction             string              `json:"faction"`
	FactionProgDetails  FactionProgDetails  `json:"faction_progression"`
}

type FactionProgDetails struct {
	GamesCount                        int     `json:"games_count"`
	IsUsedForScoreCalculation         bool    `json:"is_used_for_score_calculation"`
	RealScore                         int     `json:"real_score"`
	RealScorePeak                     int     `json:"real_score_peak"`
	UnlockedScore                     int     `json:"unlocked_score"`
	UnlockedScoreFraction             float64 `json:"unlocked_score_fraction"`
	UnlockedScoreGamesCountThreshold  int     `json:"unlocked_score_games_count_threshold"`
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
