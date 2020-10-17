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
}
