package api

func (a *apiCondition) convert() *Condition {
	return &Condition{
		Index:       a.Index,
		Name:        a.Name,
		Description: a.Desc,
	}
}

func (a *apiAbilityScore) convert() *AbilityScore {
	var skillIndices []string
	for _, skill := range a.Skills {
		skillIndices = append(skillIndices, skill.Index)
	}

	return &AbilityScore{
		Index:        a.Index,
		Name:         a.Name,
		FullName:     a.FullName,
		Description:  a.Desc,
		SkillIndices: skillIndices,
	}
}
