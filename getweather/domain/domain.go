package domain

type IdInfo struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	State   string `json:"state"`
	Country string `json:"country"`
}

type LocaleId struct {
	Status  string
	Locales []int
}

type ResgisteredCitys struct {
	Locales    []int `json:"locales"`
	MaxAllowed int   `json:"max_allowed"`
}

type WeatherInfo struct {
	Token   string
	City    string
	Weather string
}
