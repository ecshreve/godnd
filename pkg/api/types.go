package api

type Condition struct {
	Index       string
	Name        string
	Description []string
}

type DamageType struct {
	Index       string
	Name        string
	Description []string
}

type MagicSchool struct {
	Index       string
	Name        string
	Description string
}

// THIS IS WRONG
type WeaponProperty struct {
	Index       int32
	Name        string
	Description []string
}

type AbilityScore struct {
	Index        string
	Name         string
	FullName     string
	Description  []string
	SkillIndices []string
}
