package models

type ScheduleItem struct {
	Activity string `json:"activity"`
	Minutes  int    `json:"minutes"`
	Seconds  int    `json:"seconds"`
}
