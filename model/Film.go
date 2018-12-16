package model

type Film struct {
	Title         string   `json:"title"`
	Episode_id    int      `json:"episode_id"`
	Opening_crawl string   `json:"opening_crawl"`
	Director      string   `json:"director"`
	Producer      string   `json:"producer"`
	Release_date  string   `json:"release_date"`
	Characters    []string `json:"characters"`
	Planets       []string `json:"planets"`
	Starships     []string `json:"starships"`
	Vehicles      []string `json:"vehicles"`
	Species       []string `json:"species"`
	Created       string   `json:"created"`
	Edited        string   `json:"edited"`
	Url           string   `json:"url"`
}
