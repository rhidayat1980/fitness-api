package models

import "time"

type Measurements struct {
	Id         int       `json:"id"`
	UserId     int       `json:"user_id"`
	Weight     float64   `json:"weight"`
	Height     float64   `json:"height"`
	BodyFat    float64   `json:"body_fat"`
	Created_at time.Time `json:"created_at"`
}
