package dto

import "time"

type Image struct {
	Hash    string    `json:"hash" yaml:"hash" db:"hash"`
	Tittle  string    `json:"tittle" yaml:"tittle" db:"tittle"`
	Data    []byte    `json:"data" yaml:"data" db:"data"`
	AddedAt time.Time `json:"added_at" yaml:"added_at" db:"added_at"`
}
