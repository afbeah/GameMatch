package models

import()

type Play struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Stadium string `json:"stadium"`
	Date string `json:"date"`
	Result string `json:"result"`
}