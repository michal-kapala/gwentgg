package enums

type CardGroup string

const (
	Bronze CardGroup = "Bronze"
	Gold   CardGroup = "Gold"
)

type CardType string

const (
	Artifact  CardType = "artifact"
	Leader    CardType = "leader"
	Special   CardType = "special"
	Stratagem CardType = "stratagem"
	Unit      CardType = "unit"
)

type CardSet string

const (
	BaseSet    CardSet = "BaseSet"
	StarterSet CardSet = "StarterSet"
	Campaign1  CardSet = "Campaign1"
	Expansion1 CardSet = "Expansion1"
	Expansion2 CardSet = "Expansion2"
	Expansion3 CardSet = "Expansion3"
	Expansion4 CardSet = "Expansion4"
	Expansion5 CardSet = "Expansion5"
	Expansion6 CardSet = "Expansion6"
	Expansion7 CardSet = "Expansion7"
	Expansion8 CardSet = "Expansion8"
	Expansion9 CardSet = "Expansion9"
	NonOwnable CardSet = "NonOwnable"
)

type CardRarity string

const (
	Common    CardRarity = "Common"
	Rare      CardRarity = "Rare"
	Epic      CardRarity = "Epic"
	Legendary CardRarity = "Legendary"
)

type CardAction string

const (
	ArmorAdded   CardAction = "ArmorAdded"
	ArmorRemoved CardAction = "ArmorRemoved"
	Boost        CardAction = "Boost"
	Damage       CardAction = "Damage"
	Destroy      CardAction = "Destroy"
	Heal         CardAction = "Heal"
	Keyword      CardAction = "Keyword"
	Kill         CardAction = "Kill"
	Order        CardAction = "Order"
	Play         CardAction = "Play"
	Spawn        CardAction = "Spawn"
	Strengthen   CardAction = "Strengthen"
)

type CardActionSource string

const (
	Card       CardActionSource = "Card"
	CardStatus CardActionSource = "CardStatus"
	Location   CardActionSource = "Location"
)
