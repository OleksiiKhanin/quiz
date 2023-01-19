package dto

import "time"

type Card struct {
	ID          int64     `json:"id" yaml:"id"`
	Value       string    `json:"value" yaml:"value"`
	Description string    `json:"description" yaml:"description"`
	Lang        Language  `json:"lang" yaml:"lang"`
	AddedAt     time.Time `json:"added_at" yaml:"added_at"`

	Image Image `json:"image,omitempty" yaml:"image,omitempty"`
}
