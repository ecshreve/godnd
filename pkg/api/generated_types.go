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

// generated response type for api/proficiencies/{index}
type apiProficiency struct {
	Index      string         `json:"index"`
	Type       string         `json:"type"`
	Name       string         `json:"name"`
	Classes    []APIReference `json:"classes"`
	Races      []APIReference `json:"races"`
	Url        string         `json:"url"`
	References []APIReference `json:"references"`
}

// generated response type for api/languages/{index}
type apiLanguage struct {
	Index           string   `json:"index"`
	Name            string   `json:"name"`
	Type            string   `json:"type"`
	TypicalSpeakers []string `json:"typical_speakers"`
	Script          string   `json:"script"`
	Url             string   `json:"url"`
}

// generated response type for api/classes/{index}
type apiClass struct {
	Index              string         `json:"index"`
	Name               string         `json:"name"`
	HitDie             int32          `json:"hit_die"`
	ProficiencyChoices []Choice       `json:"proficiency_choices"`
	Proficiencies      []APIReference `json:"proficiencies"`
	SavingThrows       []APIReference `json:"saving_throws"`
	StartingEquipment  URLRefString   `json:"starting_equipment"`
	ClassLevels        URLRefString   `json:"class_levels"`
	Subclasses         []APIReference `json:"subclasses"`
	Spellcasting       URLRefString   `json:"spellcasting"`
	Spells             URLRefString   `json:"spells"`
	Url                string         `json:"url"`
}

// generated response type for api/classes/{index}/subclasses/
type apiSubclassesForClass struct {
	Count   int32          `json:"count"`
	Results []APIReference `json:"results"`
}

// generated response type for api/classes/{index}/spells/
type apiSpellsForClass struct {
	Count   int32          `json:"count"`
	Results []APIReference `json:"results"`
}

// generated response type for api/classes/{index}/features/
type apiFeaturesForClass struct {
	Count   int32          `json:"count"`
	Results []APIReference `json:"results"`
}

// generated response type for api/classes/{index}/proficiencies/
type apiProficienciesForClass struct {
	Count   int32          `json:"count"`
	Results []APIReference `json:"results"`
}

// generated response type for api/classes/{index}/levels/
type apiLevelsForClass struct {
	Index               string                 `json:"index"`
	Level               int32                  `json:"level"`
	AbilityScoreBonuses int32                  `json:"ability_score_bonuses"`
	ProfBonus           int32                  `json:"prof_bonus"`
	FeatureChoices      []APIReference         `json:"feature_choices"`
	Features            []APIReference         `json:"features"`
	Spellcasting        map[string]interface{} `json:"spellcasting"`
	ClassSpecific       map[string]interface{} `json:"class_specific"`
	Class               APIReference           `json:"class"`
	Url                 string                 `json:"url"`
}

// generated response type for api/classes/{index}/levels/{integer 1-20}
type apiLevelForClass struct {
	Index               string                 `json:"index"`
	Level               int32                  `json:"level"`
	AbilityScoreBonuses int32                  `json:"ability_score_bonuses"`
	ProfBonus           int32                  `json:"prof_bonus"`
	FeatureChoices      []APIReference         `json:"feature_choices"`
	Features            []APIReference         `json:"features"`
	Spellcasting        map[string]interface{} `json:"spellcasting"`
	ClassSpecific       map[string]interface{} `json:"class_specific"`
	Class               APIReference           `json:"class"`
	Url                 string                 `json:"url"`
}

// generated response type for api/classes/{index}/levels/{integer 1-20}/features
type apiFeaturesForClassAndLevel struct {
	Count   int32          `json:"count"`
	Results []APIReference `json:"results"`
}

// generated response type for api/classes/{index}/levels/{integer 1-20}/spells
type apiSpellsForClassAndLevel struct {
	Count   int32          `json:"count"`
	Results []APIReference `json:"results"`
}

// generated response type for api/conditions/{index}
type apiCondition struct {
	Index string   `json:"index"`
	Name  string   `json:"name"`
	Desc  []string `json:"desc"`
	Url   string   `json:"url"`
}

// generated response type for api/damage-types/{index}
type apiDamageType struct {
	Index string   `json:"index"`
	Name  string   `json:"name"`
	Desc  []string `json:"desc"`
	Url   string   `json:"url"`
}

// generated response type for api/magic-schools/{index}
type apiMagicSchool struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Url   string `json:"url"`
}

// generated response type for api/equipment-categories/{index}
type apiEquipmentCategory struct {
	Index     string         `json:"index"`
	Name      string         `json:"name"`
	Equipment []APIReference `json:"equipment"`
}

