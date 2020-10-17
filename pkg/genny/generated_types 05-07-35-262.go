// +build ignore

package genny

// AbilityScore is a AbilityScore
type AbilityScore struct {
	Desc     []string
	FullName string
	Index    string
	Name     string
	Skills   []Skill
	URL      string
}

// Class is a Class
type Class struct {
	Levels             string
	HitDie             int32
	Index              string
	Name               string
	Proficiencies      []Proficiency
	ProficiencyChoices []Choice
	SavingThrows       []SavingThrow
	Spellcasting       string
	Spells             string
	StartingEquipment  string
	Subclasses         []Subclass
	URL                string
}

// Condition is a Condition
type Condition struct {
	Desc  []string
	Index string
	Name  string
	URL   string
}

// DamageType is a DamageType
type DamageType struct {
	Desc  []string
	Index string
	Name  string
	URL   string
}

// Equipment is a Equipment
type Equipment struct {
	HDamage             HDamage
	ArmorCategory       string
	ArmorClass          ArmorClass
	Capacity            string
	CategoryRange       string
	Contents            []Content
	Cost                Cost
	Damage              Damage
	Desc                []string
	Category            EquipmentCategory
	GearCategory        GearCategory
	Index               string
	Name                string
	Properties          []Property
	Quantity            int32
	Range               Range
	Special             []string
	Speed               Speed
	StealthDisadvantage boolean
	StrMinimum          int32
	ThrowRange          ThrowRange
	ToolCategory        string
	URL                 string
	VehicleCategory     string
	WeaponCategory      string
	WeaponRange         string
	Weight              int32
}

// EquipmentArmorClass is a EquipmentArmorClass
type EquipmentArmorClass struct {
	Base     int32
	DexBonus boolean
	MaxBonus int32
}

// EquipmentCategory is a EquipmentCategory
type EquipmentCategory struct {
	Equipment []Equipment
	Index     string
	Name      string
	URL       string
}

// EquipmentContent is a EquipmentContent
type EquipmentContent struct {
	Item     Item
	Quantity int32
}

// EquipmentCost is a EquipmentCost
type EquipmentCost struct {
	Quantity int32
	Unit     string
}

// EquipmentDamage is a EquipmentDamage
type EquipmentDamage struct {
	DamageDice string
	DamageType DamageType
}

// EquipmentHDamage is a EquipmentHDamage
type EquipmentHDamage struct {
	DamageDice string
	DamageType DamageType
}

// EquipmentRange is a EquipmentRange
type EquipmentRange struct {
	Long   int32
	Normal int32
}

// EquipmentSpeed is a EquipmentSpeed
type EquipmentSpeed struct {
	Quantity int32
	Unit     string
}

// EquipmentThrowRange is a EquipmentThrowRange
type EquipmentThrowRange struct {
	Long   int32
	Normal int32
}

// Feature is a Feature
type Feature struct {
	Choice        Choice
	Class         Class
	Desc          []string
	Group         string
	Index         string
	Level         int32
	Name          string
	Prerequisites []Prerequisite
	Reference     string
	Subclass      Subclass
	URL           string
}

// FeaturePrerequisite is a FeaturePrerequisite
type FeaturePrerequisite struct {
	Feature string
	Level   int32
	Type    string
}

// Language is a Language
type Language struct {
	Desc            string
	Index           string
	Name            string
	Script          string
	Type            string
	TypicalSpeakers []string
	URL             string
}

// Level is a Level
type Level struct {
	AbilityScoreBonuses int32
	Class               Class
	ClassSpecific       ClassSpecific
	FeatureChoices      []Choice
	Features            []Feature
	Index               string
	Level               int32
	ProfBonus           int32
	Spellcasting        Spellcasting
	Subclass            Subclass
	SubclassSpecific    SubclassSpecific
	URL                 string
}

// LevelClassSpecific is a LevelClassSpecific
type LevelClassSpecific struct {
	ActionSurges           int32
	ArcaneRecoveryLevels   int32
	AuraRange              int32
	BardicInspirationDie   int32
	BrutalCriticalDice     int32
	ChannelDivinityCharges int32
	CreatingSpellSlots     []CreatingSpell_slot
	DestroyUndeadCr        int32
	ExtraAttacks           int32
	FavoredEnemies         int32
	FavoredTerrain         int32
	IndomitableUses        int32
	InvocationsKnown       int32
	KiPoints               int32
	MagicalSecretsMax5     int32
	MagicalSecretsMax7     int32
	MagicalSecretsMax9     int32
	MartialArts            MartialArt
	MetamagicKnown         int32
	MysticArcanumLevel6    int32
	MysticArcanumLevel7    int32
	MysticArcanumLevel8    int32
	MysticArcanumLevel9    int32
	RageCount              int32
	RageDamageBonus        int32
	SneakAttack            SneakAttack
	SongOfRestDie          int32
	SorceryPoints          int32
	UnarmoredMovement      int32
	WildShapeFly           boolean
	WildShapeMaxCr         int32
	WildShapeSwim          boolean
}

// LevelClassSpecificCreating_SpellSlot is a LevelClassSpecificCreating_SpellSlot
type LevelClassSpecificCreating_SpellSlot struct {
	SorceryPointCost int32
	SpellSlotLevel   int32
}

// LevelClassSpecificMartial_art is a LevelClassSpecificMartial_art
type LevelClassSpecificMartial_art struct {
	DiceCount int32
	DiceValue int32
}

// LevelClassSpecificSneak_attack is a LevelClassSpecificSneak_attack
type LevelClassSpecificSneak_attack struct {
	DiceCount int32
	DiceValue int32
}

// LevelSpellcasting is a LevelSpellcasting
type LevelSpellcasting struct {
	CantripsKnown    int32
	SpellSlotsLevel1 int32
	SpellSlotsLevel2 int32
	SpellSlotsLevel3 int32
	SpellSlotsLevel4 int32
	SpellSlotsLevel5 int32
	SpellSlotsLevel6 int32
	SpellSlotsLevel7 int32
	SpellSlotsLevel8 int32
	SpellSlotsLevel9 int32
	SpellsKnown      int32
}

// LevelSubclassSpecific is a LevelSubclassSpecific
type LevelSubclassSpecific struct {
	AdditionalMagicalSecretsMaxLvl int32
	AuraRange                      int32
}

// MagicItem is a MagicItem
type MagicItem struct {
	Desc              []string
	EquipmentCategory EquipmentCategory
	Index             string
	Name              string
	URL               string
}

// MagicSchool is a MagicSchool
type MagicSchool struct {
	Desc  string
	Index string
	Name  string
	URL   string
}

// Monster is a Monster
type Monster struct {
	Actions               []Action
	Alignment             string
	ArmorClass            int32
	ChallengeRating       int32
	Charisma              int32
	ConditionImmunities   []ConditionImmunity
	Constitution          int32
	DamageImmunities      []string
	DamageResistances     []string
	DamageVulnerabilities []string
	Dexterity             int32
	HitDice               string
	HitPoints             int32
	Index                 string
	Intelligence          int32
	Languages             string
	LegendaryActions      []LegendaryAction
	Name                  string
	OtherSpeeds           []OtherSpeed
	Proficiencies         []Proficiency
	Reactions             []Reaction
	Senses                Sense
	Size                  string
	SpecialAbilities      []SpecialAbility
	Speed                 Speed
	Strength              int32
	Subtype               string
	Type                  string
	URL                   string
	Wisdom                int32
}

// MonsterAction is a MonsterAction
type MonsterAction struct {
	AttackBonus   int32
	AttackOptions AttackOption
	Damage        []Damage
	Desc          string
	Name          string
	Options       Option
	Usage         Usage
}

// MonsterActionAttackOptionFrom is a MonsterActionAttackOptionFrom
type MonsterActionAttackOptionFrom struct {
	Damage []Damage
	Dc     Dc
	Name   string
}

// MonsterActionAttackOptionFromDamage is a MonsterActionAttackOptionFromDamage
type MonsterActionAttackOptionFromDamage struct {
	DamageDice string
	DamageType DamageType
}

// MonsterActionAttackOptionFromDc is a MonsterActionAttackOptionFromDc
type MonsterActionAttackOptionFromDc struct {
	DcType      DcType
	DcValue     int32
	SuccessType string
}

// MonsterActionDamage is a MonsterActionDamage
type MonsterActionDamage struct {
	DamageDice string
	DamageType DamageType
}

// MonsterActionOption is a MonsterActionOption
type MonsterActionOption struct {
	Choose int32
}

// MonsterActionUsage is a MonsterActionUsage
type MonsterActionUsage struct {
	Dice     string
	MinValue int32
	Type     string
}

// MonsterLegendaryAction is a MonsterLegendaryAction
type MonsterLegendaryAction struct {
	Desc string
	Name string
}

// MonsterOtherSpeed is a MonsterOtherSpeed
type MonsterOtherSpeed struct {
	Form  string
	Speed Speed
}

// MonsterOtherSpeedSpeed is a MonsterOtherSpeedSpeed
type MonsterOtherSpeedSpeed struct {
	Climb string
	Walk  string
}

// MonsterProficiency is a MonsterProficiency
type MonsterProficiency struct {
	Proficiency Proficiency
	Value       int32
}

// MonsterReaction is a MonsterReaction
type MonsterReaction struct {
	Desc string
	Name string
}

// MonsterSense is a MonsterSense
type MonsterSense struct {
	Blindsight        string
	Darkvision        string
	PassivePerception int32
	Tremorsense       string
	Truesight         string
}

// MonsterSpecialAbility is a MonsterSpecialAbility
type MonsterSpecialAbility struct {
	Desc string
	Name string
}

// MonsterSpeed is a MonsterSpeed
type MonsterSpeed struct {
	Burrow string
	Climb  string
	Fly    string
	Hover  boolean
	Swim   string
	Walk   string
}

// Proficiency is a Proficiency
type Proficiency struct {
	Classes    []Class
	Index      string
	Name       string
	Races      []Race
	References []Reference
	Type       string
	URL        string
}

// ProficiencyReference is a ProficiencyReference
type ProficiencyReference struct {
	Index string
	Name  string
	Type  string
	URL   string
}

// Race is a Race
type Race struct {
	AbilityBonusOptions        AbilityBonus_option
	AbilityBonuses             []AbilityBonuse
	Age                        string
	Alignment                  string
	Index                      string
	LanguageDesc               string
	LanguageOptions            LanguageOption
	Languages                  []Language
	Name                       string
	Size                       string
	SizeDescription            string
	Speed                      int32
	StartingProficiencies      []StartingProficiency
	StartingProficiencyOptions StartingProficiency_option
	Subraces                   []Subrace
	TraitOptions               TraitOption
	Traits                     []Trait
	URL                        string
}

// RaceAbilityBonus_optionFrom is a RaceAbilityBonus_optionFrom
type RaceAbilityBonus_optionFrom struct {
	Bonus int32
	Index string
	Name  string
	URL   string
}

// RaceAbilityBonuse is a RaceAbilityBonuse
type RaceAbilityBonuse struct {
	Bonus int32
	Index string
	Name  string
	URL   string
}

// Rule is a Rule
type Rule struct {
	Desc        string
	Index       string
	Name        string
	Subsections []Subsection
	URL         string
}

// RuleSection is a RuleSection
type RuleSection struct {
	Desc  string
	Index string
	Name  string
	URL   string
}

// Skill is a Skill
type Skill struct {
	AbilityScore AbilityScore
	Desc         []string
	Index        string
	Name         string
	URL          string
}

// Spell is a Spell
type Spell struct {
	AreaOfEffect    AreaOf_effect
	AttackType      string
	CastingTime     string
	Classes         []Class
	Components      []string
	Concentration   boolean
	Damage          Damage
	Dc              Dc
	Desc            []string
	Duration        string
	HealAtSlotLevel HealAt_SlotLevel
	HigherLevel     []string
	Index           string
	Level           int32
	Material        string
	Name            string
	Range           string
	Ritual          boolean
	School          School
	Subclasses      []Subclass
	URL             string
}

// SpellAreaOf_effect is a SpellAreaOf_effect
type SpellAreaOf_effect struct {
	Size int32
	Type string
}

// Spellcasting is a Spellcasting
type Spellcasting struct {
	Class   Class
	Index   string
	Info    []Info
	Level   int32
	Ability SpellcastingAbility
	URL     string
}

// SpellcastingInfo is a SpellcastingInfo
type SpellcastingInfo struct {
	Desc []string
	Name string
}

// SpellDamage is a SpellDamage
type SpellDamage struct {
	DamageAtCharacterLevel DamageAt_CharacterLevel
	DamageAtSlotLevel      DamageAt_SlotLevel
	DamageType             DamageType
}

// SpellDamageDamageAt_CharacterLevel is a SpellDamageDamageAt_CharacterLevel
type SpellDamageDamageAt_CharacterLevel struct {
	A1  string
	A11 string
	A17 string
	A5  string
}

// SpellDamageDamageAt_SlotLevel is a SpellDamageDamageAt_SlotLevel
type SpellDamageDamageAt_SlotLevel struct {
	A1 string
	A2 string
	A3 string
	A4 string
	A5 string
	A6 string
	A7 string
	A8 string
	A9 string
}

// SpellDc is a SpellDc
type SpellDc struct {
	DcSuccess string
	DcType    DcType
	Desc      string
}

// SpellHealAt_SlotLevel is a SpellHealAt_SlotLevel
type SpellHealAt_SlotLevel struct {
	A1 string
	A2 string
	A3 string
	A4 string
	A5 string
	A6 string
	A7 string
	A8 string
	A9 string
}

// Startingequipment is a Startingequipment
type Startingequipment struct {
	Class                    Class
	Index                    string
	StartingEquipment        []StartingEquipment
	StartingEquipmentOptions []StartingEquipment_option
	URL                      string
}

// StartingequipmentStartingEquipment is a StartingequipmentStartingEquipment
type StartingequipmentStartingEquipment struct {
	Equipment Equipment
	Quantity  int32
}

// StartingequipmentStartingEquipment_optionFrom is a StartingequipmentStartingEquipment_optionFrom
type StartingequipmentStartingEquipment_optionFrom struct {
	Equipment Equipment
	Quantity  int32
}

// Subclass is a Subclass
type Subclass struct {
	Class  Class
	Desc   []string
	Index  string
	Name   string
	Spells []Spell
	Flavor string
	Levels string
	URL    string
}

// SubclassSpell is a SubclassSpell
type SubclassSpell struct {
	Prerequisites []Prerequisite
	Spell         Spell
}

// SubclassSpellPrerequisite is a SubclassSpellPrerequisite
type SubclassSpellPrerequisite struct {
	Index string
	Name  string
	Type  string
	URL   string
}

// Subrace is a Subrace
type Subrace struct {
	AbilityBonuses        []AbilityBonuse
	Desc                  string
	Index                 string
	LanguageOptions       LanguageOption
	Name                  string
	Race                  Race
	RacialTraitOptions    RacialTrait_option
	RacialTraits          []RacialTrait
	StartingProficiencies []StartingProficiency
	URL                   string
}

// SubraceAbilityBonuse is a SubraceAbilityBonuse
type SubraceAbilityBonuse struct {
	Bonus int32
	Index string
	Name  string
	URL   string
}

// Trait is a Trait
type Trait struct {
	Desc               []string
	Index              string
	Name               string
	Proficiencies      []Proficiency
	ProficiencyChoices ProficiencyChoice
	Races              []Race
	Subraces           []Subrace
	URL                string
}

// WeaponProperty is a WeaponProperty
type WeaponProperty struct {
	Desc  []string
	Index string
	Name  string
	URL   string
}
