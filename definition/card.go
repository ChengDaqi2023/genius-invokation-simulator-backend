/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "card.go" LastUpdatedAt 2022/12/12 10:19:12
 */

package definition

// CardType 卡牌类型，大类
type CardType byte

// CardSubType 卡牌类型，小类
type CardSubType byte

const (
	CardEvent     CardType = iota << 4 // CardEvent 事件卡
	CardSupport                        // CardSupport 支援卡
	CardEquipment                      // CardEquipment 装备卡
)

// EventCardSubType
const (
	CardElementalResonance CardType = CardEvent + 1 + iota // CardElementalResonance 元素共鸣事件卡
	CardFood                                               // CardFood 食物事件卡
)

// SupportCardSubType
const (
	CardLocation  CardType = CardSupport + 1 + iota // CardLocation 地点支援卡
	CardCompanion                                   // CardCompanion 伙伴支援卡
	CardItem                                        // CardItem 物品支援卡
)

// EquipmentCardSubType
const (
	CardTalent   CardType = CardEquipment + 1 + iota // CardTalent 天赋装备卡
	CardWeapon                                       // CardWeapon 武器装备卡
	CardArtifact                                     // CardArtifact 圣遗物装备卡
)
