// Code generated by genny; DO NOT EDIT.

package api

// generated response type for api/ability-scores/{index}
type apiAbilityScore struct {
	Index    string         `json:"index"`
	Name     string         `json:"name"`
	FullName string         `json:"full_name"`
	Desc     []string       `json:"desc"`
	Skills   []APIReference `json:"skills"`
	Url      string         `json:"url"`
}

// generated response type for api/skills/{index}
type apiSkill struct {
	Index        string       `json:"index"`
	Name         string       `json:"name"`
	Desc         []string     `json:"desc"`
	AbilityScore APIReference `json:"ability_score"`
	Url          string       `json:"url"`
}

