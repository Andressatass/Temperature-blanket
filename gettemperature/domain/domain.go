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