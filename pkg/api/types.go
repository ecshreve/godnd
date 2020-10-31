package api

//go:generate genny

// ResourceList is a list of APIReferences for an endpoint.
type ResourceList struct {
	Count     int            `json:"count"`
	Resources []APIReference `json:"results"`
}

// APIReference is the basic data about an API resource.
type APIReference struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

// Choice is a selection of Choose number of ResourceTypes out of the the From list.
type Choice struct {
	Choose       int32          `json:"choose"`
	ResourceType string         `json:"type"`
	From         []APIReference `json:"from"`
}

// Cost is a monetary amount with a Quantity and Unit.
type Cost struct {
	Quantity int32  `json:"quantity"`
	Unit     string `json:"unit"`
}

// AbilityBonus is a bonus modifier to the associated AbilityScore.
type AbilityBonus struct {
	Bonus        int32        `json:"bonus"`
	AbilityScore APIReference `json:"ability_score"`
}
