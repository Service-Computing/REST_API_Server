package model

type Planet struct {
	Name            string   `json:"name"`
	Rotation_period string   `json:"rotation_period"`
	Orbital_period  string   `json:"orbital_period"`
	Diameter        string   `json:"diameter"`
	Climate         string   `json:"climate"`
	Gravity         string   `json:"gravity"`
	Terrain         string   `json:"terrain"`
	Surface_water   string   `json:"surface_water"`
	Population      string   `json:"population"`
	Residents       []string `json:"residents"`
	Films           []string `json:"films"`
	Created         string   `json:"created"`
	Edited          string   `json:"edited"`
	Url             string   `json:"url"`
}
