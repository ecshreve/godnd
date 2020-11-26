package api

type BasicResource struct {
	Index       string
	Name        string
	Description []string
}

type Condition BasicResource

type AbilityScore struct {
	Index        string
	Name         string
	FullName     string
	Description  []string
	SkillIndices []string
}
