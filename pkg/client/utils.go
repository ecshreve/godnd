package client

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
