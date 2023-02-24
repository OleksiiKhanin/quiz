package dto

import "time"

type Rating struct {
	CardID     int64     `json:"card_id"`
	ShowsCount int64     `json:"shows_count"`
	Stars      float32   `json:"stars"`
	ModifiedAt time.Time `json:"modified_at"`
}
