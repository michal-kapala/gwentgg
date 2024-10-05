package seasons

import (
	"gwentgg/db/models"
	"gwentgg/services"
)

type SeasonList struct {
	TotalCount    uint           `json:"total_count"`
	Limit         uint           `json:"limit"`
	PageToken     string         `json:"page_token"`
	NextPageToken *uint          `json:"next_page_token"`
	Items         []RankedSeason `json:"items"`
}

type RankedSeason struct {
	ID                      string                      `json:"id"`
	Name                    string                      `json:"name"`
	Ranks                   []Rank                      `json:"ranks"`
	DateStarts              string                      `json:"date_starts"`
	DateEnds                string                      `json:"date_ends"`
	DateCreated             string                      `json:"date_created"`
	RankingType             string                      `json:"ranking_type"`
	ParentSeason            *RankedSeason               `json:"parent_season"`
	SeasonType              *string                     `json:"season_type"`
	InitialRank             *uint                       `json:"initial_rank"`
	InitialMosaicPieceCount *uint                       `json:"initial_mosaic_piece_count"`
	LocalizationInfo        []services.LocalizationInfo `json:"localization_info"`
	IsCurrentForUser        *bool                       `json:"is_current_for_user"`
}

func (season *RankedSeason) ToModel() (models.Season, error) {
	starts, err := services.ParseDate(season.DateStarts)
	if err != nil {
		return models.Season{}, err
	}

	ends, err := services.ParseDate(season.DateEnds)
	if err != nil {
		return models.Season{}, err
	}

	created, err := services.ParseDate(season.DateCreated)
	if err != nil {
		return models.Season{}, err
	}

	var parentID *string
	if season.ParentSeason != nil {
		parentID = &season.ParentSeason.ID
	}

	s := models.Season{
		ID:           season.ID,
		Name:         season.Name,
		DateStarts:   starts,
		DateEnds:     ends,
		DateCreated:  created,
		RankingType:  season.RankingType,
		ParentSeason: parentID,
		SeasonType:   season.SeasonType,
	}
	return s, nil
}

func ToModels(seasons []RankedSeason) ([]models.Season, error) {
	var seasonModels []models.Season
	for _, s := range seasons {
		seasonModel, err := s.ToModel()
		if err != nil {
			return seasonModels, err
		}
		seasonModels = append(seasonModels, seasonModel)
	}
	return seasonModels, nil
}

type Rank struct {
	ID                        string `json:"id"`
	Name                      string `json:"name"`
	ScoreFrom                 int    `json:"score_from"`
	ScoreTo                   int    `json:"score_to"`
	MosaicPiecesRequired      uint   `json:"mosaic_pieces_required"`
	TierID                    uint   `json:"tier_id"`
	FactionScoreLowerBoundary uint   `json:"faction_score_lower_boundary"`
	FactionScoreUpperBoundary uint   `json:"faction_score_upper_boundary"`
	RankInNextSeason          uint   `json:"rank_in_next_season"`
}
