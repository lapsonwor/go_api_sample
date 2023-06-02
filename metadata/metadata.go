package metadata

type Metadata struct {
	Name      string       `json:"name"`
	Image			string				`json:"image"`
	ExternalURL string			`json:"external_url"`
	Attributes  []Attribute `json:"attributes"`
	
}

type Attribute struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}