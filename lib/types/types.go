package types

type WeatherSearcher interface {
	GetByPostCode(input string) (SearchResult, error)
	GetByLocationName(input string) (SearchResult, error)
}

type SearchResult struct {
	SearchInput string
	Name        string  `json:"name"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Country     string  `json:"country"`
	State       string  `json:"state,omitempty"`
	Error       string
}
