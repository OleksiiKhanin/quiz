package dto

import "time"

type Image struct {
	Hash    string    `json:"hash" yaml:"hash" db:"hash"`
	Title   string    `json:"title" yaml:"title" db:"title"`
	Type    string    `json:"type" yaml:"type" db:"type"`
	Data    []byte    `json:"data" yaml:"data" db:"data"`
	AddedAt time.Time `json:"added_at" yaml:"added_at" db:"added_at"`
}
