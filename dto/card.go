package dto

import "time"

type Card struct {
	ID          int64     `json:"id" yaml:"id" db:"id"`
	Value       string    `json:"value" yaml:"value" db:"value"`
	Description string    `json:"description" yaml:"description" db:"description"`
	Lang        Language  `json:"lang" yaml:"lang" db:"lang"`
	AddedAt     time.Time `json:"added_at" yaml:"added_at" db:"added_at"`
	ImageHash   string    `json:"image_hash" yaml:"image_hash" db:"image_hash"`
}

type CreateCardsRequest struct {
	Cards       [2]*Card `json:"cards"`
	ImageData   []byte   `json:"image_data"`
	ImageTittle string   `json:"image_title"`
	ImageType   string   `json:"image_type"`
}
