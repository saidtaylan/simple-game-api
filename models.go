package main

type Game struct {
	ID            int
	Name          string `json:"name" validate:"required"`
	LatestVersion string `json:"latest_version" validate:"required"`
	Category      string `json:"category" validate:"required"`
	Descr         string `json:"descr" validate:"required"`
	Producer      string `json:"producer" validate:"required"`
}
