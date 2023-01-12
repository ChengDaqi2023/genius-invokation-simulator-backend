package entity

import (
	"github.com/sunist-c/genius-invokation-simulator-backend/enum"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/context"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/event"
	"github.com/sunist-c/genius-invokation-simulator-backend/model/modifier"
)

type PlayerInfo struct {
	UID        uint
	Cards      []Card
	Characters []*character
}

type player struct {
	uid    uint              // uid 玩家的UID，由其他模块托管
	status enum.PlayerStatus // status 玩家的状态

	operated    bool  // operated 本回合玩家是否操作过
	reRollTimes uint  // reRollTimes 重新投掷的次数
	staticCost  *Cost // staticCost 每回合投掷阶段固定产出的骰子

	holdingCost     *Cost         // holdingCost 玩家持有的骰子
	cardDeck        CardDeck      // cardDeck 玩家的牌堆
	holdingCards    map[uint]Card // holdingCards 玩家持有的卡牌
	activeCharacter uint          // activeCharacter 玩家当前的前台角色

	characters    map[uint]*character // characters 玩家出战的角色
	characterList []uint              // characterList 玩家的角色列表
	summons       map[uint]Summon     // summons 玩家在场的召唤物
	summonList    []uint              // summonList 玩家的召唤物列表
	supports      map[uint]Support    // supports 玩家在场的支援
	supportList   []uint              // supportList 玩家的支援物列表

	globalDirectAttackModifiers AttackModifiers  // globalDirectAttackModifiers 全局直接攻击修正
	globalFinalAttackModifiers  AttackModifiers  // globalFinalAttackModifiers 全局最终攻击修正
	globalDefenceModifiers      DefenceModifiers // globalDefenceModifiers 全局防御修正
	globalHealModifiers         HealModifiers    // globalHealModifiers 全局治疗修正
	globalChargeModifiers       ChargeModifiers  // globalChargeModifiers 全局充能修正
	globalCostModifiers         CostModifiers    // globalCostModifiers 全局费用修正

	cooperativeAttacks []CooperativeSkill // cooperativeAttacks 协同攻击技能
	callbackEvents     *event.Map         // callbackEvents 回调事件集合
}

func (p player) GetUID() (uid uint) {
	return p.uid
}

func (p player) GetCost() (cost map[enum.ElementType]uint) {
	return p.holdingCost.costs
}

func (p player) GetCards() (cards []uint) {
	cards = []uint{}
	for i := range p.holdingCards {
		cards = append(cards, i)
	}

	return cards
}

func (p player) GetSummons() (summons []uint) {
	return p.summonList
}

func (p player) GetSupports() (supports []uint) {
	return p.summonList
}

func (p player) CardDeckRemain() (remain uint) {
	return p.cardDeck.remain
}

func (p player) GetActiveCharacter() (character uint) {
	return p.activeCharacter
}

func (p player) GetBackgroundCharacters() (characters []uint) {
	index := 0
	for i, id := range p.characterList {
		if id == p.activeCharacter {
			index = i
		}
	}

	return append(p.characterList[:index], p.characterList[index+1:]...)
}

func (p player) GetCharacter(character uint) (has bool, entity Character) {
	characterEntity, exist := p.characters[character]
	return exist, characterEntity
}

func (p player) GetStatus() (status enum.PlayerStatus) {
	return p.status
}

func (p player) GetGlobalModifiers(modifierType enum.ModifierType) (modifiers []uint) {
	switch modifierType {
	case enum.ModifierTypeNone:
		return []uint{}
	case enum.ModifierTypeAttack:
		modifiers = p.globalDirectAttackModifiers.Expose()
		modifiers = append(modifiers, p.globalFinalAttackModifiers.Expose()...)
		return modifiers
	case enum.ModifierTypeCharacter:
		return []uint{}
	case enum.ModifierTypeCharge:
		return p.globalChargeModifiers.Expose()
	case enum.ModifierTypeCost:
		return p.globalCostModifiers.Expose()
	case enum.ModifierTypeDefence:
		return p.globalDefenceModifiers.Expose()
	case enum.ModifierTypeHeal:
		return p.globalHealModifiers.Expose()
	default:
		return []uint{}
	}
}

func (p player) GetCooperativeSkills(trigger enum.TriggerType) (skills []uint) {
	return []uint{}
}

func (p player) GetEvents(trigger enum.TriggerType) (events []uint) {
	return p.callbackEvents.Expose(trigger)
}

func newPlayer(info PlayerInfo) *player {
	player := &player{
		uid:                         info.UID,
		status:                      enum.PlayerStatusViewing,
		operated:                    false,
		reRollTimes:                 1,
		staticCost:                  NewCost(),
		holdingCost:                 NewCost(),
		cardDeck:                    *NewCardDeck(info.Cards),
		holdingCards:                map[uint]Card{},
		activeCharacter:             0,
		characters:                  map[uint]*character{},
		characterList:               []uint{},
		summons:                     map[uint]Summon{},
		summonList:                  []uint{},
		supports:                    map[uint]Support{},
		supportList:                 []uint{},
		globalDirectAttackModifiers: modifier.NewChain[context.DamageContext](),
		globalFinalAttackModifiers:  modifier.NewChain[context.DamageContext](),
		globalDefenceModifiers:      modifier.NewChain[context.DamageContext](),
		globalHealModifiers:         modifier.NewChain[context.HealContext](),
		globalChargeModifiers:       modifier.NewChain[context.ChargeContext](),
		globalCostModifiers:         modifier.NewChain[context.CostContext](),
		cooperativeAttacks:          []CooperativeSkill{},
		callbackEvents:              event.NewEventMap(),
	}

	// todo: complete ctor of player
	return player
}
