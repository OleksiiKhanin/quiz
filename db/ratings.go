package db

import (
	"context"
	"english-card/dto"
	"english-card/interfaces"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const ratingsTableName = "ratings"

type ratingRepo struct {
	db *sqlx.DB
}

func (r *ratingRepo) Begin(t interfaces.TransactionGetter) (*ratingTx, error) {
	if t != nil {
		if tx := t.GetTX(); tx != nil {
			return &ratingTx{tx: tx}, nil
		}
	}
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}
	return &ratingTx{tx: tx}, nil
}

type ratingTx struct {
	tx *sqlx.Tx
}

func (r *ratingTx) InsertRating(ctx context.Context, cardID int64) error {
	_, err := r.tx.ExecContext(ctx, fmt.Sprintf("INSERT INTO %s (card_id, stars, shows) VALUES (%d, 0, 1)", ratingsTableName, cardID))
	if err != nil {
		r.tx.Rollback()
	}
	return err
}

func (r *ratingTx) UpdateRating(ctx context.Context, cardID int64, star uint8) (dto.Rating, error) {
	query := fmt.Sprintf("UPDATE %s SET stars=(stars*shows+%d)/shows, modified_at='NOW()' WHERE card_id = %d RETURNING stars, shows, modified_at", ratingsTableName, star, cardID)
	rating := dto.Rating{
		CardID: cardID,
	}
	err := r.tx.QueryRowxContext(ctx, query).Scan(&rating.Stars, &rating.ShowsCount, &rating.ModifiedAt)
	if err != nil {
		r.tx.Rollback()
	}
	return rating, err
}

func (r *ratingTx) GetRating(ctx context.Context, cardID int64) (dto.Rating, error) {
	query := fmt.Sprintf("UPDATE %s SET shows = shows+1, modified_at='NOW()' WHERE card_id = %d RETURNING stars, shows, modified_at", ratingsTableName, cardID)
	rating := dto.Rating{
		CardID: cardID,
	}
	err := r.tx.QueryRowxContext(ctx, query).Scan(&rating.Stars, &rating.ShowsCount, &rating.ModifiedAt)
	if err != nil {
		r.tx.Rollback()
	}
	return rating, err
}

func (r *ratingTx) End() error {
	return r.tx.Commit()
}
func (r *ratingTx) GetTX() *sqlx.Tx {
	return r.tx
}
