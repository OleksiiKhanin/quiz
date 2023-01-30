package interfaces

import (
	"context"
	"english-card/dto"
	"github.com/jmoiron/sqlx"
)

type ImageService interface {
	GetImage(ctx context.Context, hash string) (dto.Image, error)
	SaveImage(ctx context.Context, tittle, typ string, image []byte) (string, error)
}

type CardService interface {
	AddCards(ctx context.Context, image *dto.Image, cards ...[2]*dto.Card) error
	UpdateCard(ctx context.Context, card *dto.Card) error
	GetCards(ctx context.Context, id int64) ([]dto.Card, error)
	GetRandomCards(ctx context.Context, lang dto.Language) ([]dto.Card, error)
}

type ImageRepo interface {
	Begin(t TransactionGetter) (ImageTx, error)
}

type CardRepo interface {
	Begin(t TransactionGetter) (CardTx, error)
}

type End interface {
	End() error
}

type ImageTx interface {
	TransactionGetter
	GetImage(ctx context.Context, hash string) (dto.Image, error)
	InsertImage(ctx context.Context, tittle, typ string, image []byte) (string, error)
	End
}

type CardTx interface {
	TransactionGetter
	InsertCard(ctx context.Context, card *dto.Card) (int64, error)
	UpdateCard(ctx context.Context, card *dto.Card) error
	CreatePairs(ctx context.Context, ids [2]int64) error
	GetCard(ctx context.Context, id int64) (dto.Card, error)
	GetPairsIDs(ctx context.Context, id int64) ([]int64, error)
	GetIDs(ctx context.Context, lang dto.Language, limit int64) ([]int64, error)
	End
}

type TransactionGetter interface {
	GetTX() *sqlx.Tx
}
