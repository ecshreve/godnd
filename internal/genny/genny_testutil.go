package genny

var TestGqlTypeString = map[string]string{
	"AbilityScore": `type AbilityScore {
  _id: ObjectId
  desc: [String]
  full_name: String
  index: String
  name: String
  skills: [AbilityScoreSkill]
  url: String
}`,
	"Class": `type Class {
  _id: ObjectId
  class_levels: String
  hit_die: Int
  index: String
  name: String
  proficiencies: [ClassProficiency]
  proficiency_choices: [ClassProficiency_choice]
  saving_throws: [ClassSaving_throw]
  spellcasting: String
  spells: String
  starting_equipment: String
  subclasses: [ClassSubclass]
  url: String
}`,
	"Condition": `type Condition {
  _id: ObjectId
  desc: [String]
  index: String
  name: String
  url: String
}`,
}

var TestGoTypeString = map[string]string{
	"AbilityScore": `type AbilityScore struct {
	Desc []string
	FullName string
	Index string
	Name string
	Skills []Skill
	URL string
}`,
	"Class": `type Class struct {
	Levels string
	HitDie int32
	Index string
	Name string
	Proficiencies []Proficiency
	ProficiencyChoices []Proficiency_choice
	SavingThrows []Saving_throw
	Spellcasting string
	Spells string
	StartingEquipment string
	Subclasses []Subclass
	URL string
}`,
	"Condition": `type Condition struct {
	Desc []string
	Index string
	Name string
	URL string
}`,
}
