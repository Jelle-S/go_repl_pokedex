package models

type CliCommand struct {
	Name        string
	Description string
	Callback    func(config *ConfigType) error
}

type ConfigType struct {
	Next     string
	Previous string
}

type LocationAreaResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
