package cards

import (
	"fmt"
	"gwentgg/db/models"
	"gwentgg/enums"
	"gwentgg/services"
	"gwentgg/utils"
)

type CardDefList struct {
	TotalCount    uint             `json:"total_count"`
	Limit         uint             `json:"limit"`
	PageToken     string           `json:"page_token"`
	NextPageToken *uint            `json:"next_page_token"`
	Items         []CardDefinition `json:"items"`
}

type CardDefinition struct {
	ID                       string                      `json:"id"`
	Template                 CardTemplate                `json:"card_template"`
	ArtID                    string                      `json:"art_id"`
	Premium                  bool                        `json:"premium"`
	Availability             string                      `json:"availability"`
	Rarity                   string                      `json:"rarity"`
	DateCreated              string                      `json:"date_created"`
	IsDeleted                bool                        `json:"is_deleted"`
	MillingProfitDust        *uint                       `json:"milling_profit_dust"`
	CraftingCostDust         *uint                       `json:"crafting_cost_dust"`
	TransmutationCostPowder  *uint                       `json:"transmutation_cost_powder"`
	MillingProfitPowder      *uint                       `json:"milling_profit_powder"`
	ReplacedByCardDefinition *string                     `json:"replaced_by_card_definition"`
	LocalizationInfo         []services.LocalizationInfo `json:"localization_info"`
	ImageAssets              []services.ImageAsset       `json:"image_assets"`
}

func (def CardDefinition) ToModel() (models.CardDefinition, error) {
	created, err := utils.ParseDate(def.DateCreated)
	if err != nil {
		return models.CardDefinition{}, err
	}

	cats := ""
	for _, cat := range def.Template.CategoryIDs {
		if len(cats) > 0 {
			cats += ","
		}
		cats += fmt.Sprintf("%d", cat)
	}

	factions := ""
	for _, faction := range def.Template.SecondaryFactions {
		if len(factions) > 0 {
			factions += ","
		}
		factions += faction
	}

	cardDef := models.CardDefinition{
		ID:                def.ID,
		TemplateID:        def.Template.ID,
		ArtID:             def.ArtID,
		Name:              def.Template.Name,
		Faction:           enums.Faction(def.Template.Faction),
		Power:             def.Template.Power,
		CardGroup:         enums.CardGroup(def.Template.CardGroup),
		ProvisionCost:     def.Template.ProvisionsCost,
		Type:              enums.CardType(def.Template.Type),
		PrimaryCategoryID: def.Template.PrimaryCategoryID,
		CategoryIDs:       cats,
		SecondaryFactions: factions,
		Armor:             def.Template.Armour,
		Premium:           def.Premium,
		Availability:      enums.CardSet(def.Availability),
		Rarity:            enums.CardRarity(def.Rarity),
		DateCreated:       created,
		IsDeleted:         def.IsDeleted,
	}
	return cardDef, err
}

func ToModels(defs []CardDefinition) ([]models.CardDefinition, error) {
	var cardDefModels []models.CardDefinition
	for _, def := range defs {
		defModel, err := def.ToModel()
		if err != nil {
			return cardDefModels, err
		}
		cardDefModels = append(cardDefModels, defModel)
	}
	return cardDefModels, nil
}

type CardTemplate struct {
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	Faction           string   `json:"faction"`
	Power             uint     `json:"power"`
	CardGroup         string   `json:"card_group"`
	DateCreated       string   `json:"date_created"`
	ProvisionsCost    uint     `json:"provisions_cost"`
	Type              string   `json:"type"`
	PrimaryCategoryID uint     `json:"primary_category_id"`
	CategoryIDs       []uint   `json:"category_ids"`
	SecondaryFactions []string `json:"secondary_factions"`
	Armour            uint     `json:"armour"`
}
