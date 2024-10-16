package model

type Match struct{
	ID int `json:"id"`
	Sport string `json:"sport"`
	Date string `json:"date"`
	ChampionShip string `json:"championship"`
	HomeTeam string `json:"hometeam"`
	VisitingTeam string `json:"visitingteam"`
	Stadium string `json:"stadium"`
	ScoreBoard string `json:"scoreboard"`
	Note string `json:"note"`
}