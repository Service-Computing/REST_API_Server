package model

type People struct {
	Name       string   `json:"name"`
	Height     string   `json:"height"`
	Mass       string   `json:"mass"`
	Hair_color string   `json:"hair_color"`
	Skin_color string   `json:"skin_color"`
	Eye_color  string   `json:"eye_color"`
	Birth_year string   `json:"birth_year"`
	Gender     string   `json:"gender"`
	Homeworld  string   `json:"homeworld"`
	Films      []string `json:"films"`
	Species    []string `json:"species"`
	Vehicles   []string `json:"vehicles"`
	Starships  []string `json:"starships"`
	Created    string   `json:"created"`
	Edited     string   `json:"edited"`
	Url        string   `json:"url"`
}

type Peoples []People
