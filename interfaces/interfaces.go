package interfaces

import (
	"context"
	"english-card/dto"
)

type ImageService interface {
	GetImage(ctx context.Context, hash string) (dto.Image, error)
	SaveImage(ctx context.Context, tittle string, image []byte) (string, error)
}

type CardService interface {
	AddCards(ctx context.Context, cards ...[2]*dto.Card) error
	GetCards(ctx context.Context, id int64) ([]dto.Card, error)
	GetRandomCards(ctx context.Context, lang dto.Language) ([]dto.Card, error)
}

type ImageRepo interface {
	Begin() (ImageTx, error)
}

type CardRepo interface {
	Begin() (CardTx, error)
}

type ImageTx interface {
	GetImage(ctx context.Context, hash string) (dto.Image, error)
	SaveImage(ctx context.Context, tittle string, image []byte) (string, error)
	End() error
}

type CardTx interface {
	InsertCard(ctx context.Context, card *dto.Card) (int64, error)
	CreateComplience(ctx context.Context, ids [2]int64) error
	GetCard(ctx context.Context, id int64) (dto.Card, error)
	GetComplianceCards(ctx context.Context, id int64) ([]dto.Card, error)
	GetIDs(ctx context.Context, lang dto.Language, limit int64) ([]int64, error)
	End() error
}
