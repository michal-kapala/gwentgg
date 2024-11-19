package models

import (
	"gwentgg/enums"
	"gwentgg/utils"

	"gorm.io/gorm"
)

type CardAction struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey;autoIncrement:false"`
	PlayerID    string `gorm:"primaryKey"`
	GameID      string `gorm:"primaryKey"`
	Action      enums.CardAction
	Amount      uint
	SourceType  enums.CardActionSource
	SourceID    uint
	SourceOwner string
	TargetID    *uint
	TargetOwner *string
	RoundID     uint
	TurnID      uint
}

func (action CardAction) FindSource(defs []CardDefinition) *CardDefinition {
	var result CardDefinition
	srcID := utils.ToString(action.SourceID)
	for _, def := range defs {
		if srcID == def.ID {
			return &def
		}
	}
	return &result
}

func (action CardAction) FindTarget(defs []CardDefinition) *CardDefinition {
	var result CardDefinition
	if action.TargetID != nil {
		targetID := utils.ToString(*action.TargetID)
		for _, def := range defs {
			if targetID == def.ID {
				return &def
			}
		}
	}
	return &result
}

type CardActionView struct {
	Action *CardAction
	Source *CardDefinition
	Target *CardDefinition
	TurnOf string
}

func MaxTurn(actions []CardActionView) uint {
	if len(actions) == 0 {
		return 0
	}
	max := uint(1)
	for _, action := range actions {
		if action.Action.TurnID > max {
			max = action.Action.TurnID
		}
	}
	return max
}

func FilterTurn(actions []CardActionView, turnID uint) []CardActionView {
	result := []CardActionView{}
	for _, action := range actions {
		if action.Action.TurnID == turnID {
			result = append(result, action)
		}
	}
	return result
}

type GameCardActions struct {
	Round1 []CardActionView
	Round2 []CardActionView
	Round3 []CardActionView
}
