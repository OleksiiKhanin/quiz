package db

import (
	"context"
	"english-card/dto"
	"english-card/interfaces"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const cardTableName = "cards"
const pairsTableName = "pairs"

func GetCardRepo(db *sqlx.DB) interfaces.CardRepo {
	return &cardRepo{db: db}
}

type cardRepo struct {
	db *sqlx.DB
}

func (c *cardRepo) Begin(t interfaces.TransactionGetter) (interfaces.CardTx, error) {
	if t != nil {
		if tx := t.GetTX(); tx != nil {
			return &cardTransaction{tx: tx}, nil
		}
	}
	tx, err := c.db.Beginx()
	if err != nil {
		return nil, err
	}
	return &cardTransaction{tx: tx}, nil
}

type cardTransaction struct {
	tx *sqlx.Tx
}

func (c *cardTransaction) GetTX() *sqlx.Tx {
	return c.tx
}

func (c *cardTransaction) InsertCard(ctx context.Context, card *dto.Card) (int64, error) {
	var id int64
	query := fmt.Sprintf("INSERT INTO %s (value, description, lang, image_hash) VALUES ($1,$2,$3,$4) RETURNING id", cardTableName)
	err := c.tx.QueryRowxContext(
		ctx,
		query,
		card.Value,
		card.Description,
		card.Lang,
		card.ImageHash,
	).Scan(&id)
	if err != nil {
		c.tx.Rollback()
		return 0, err
	}
	return id, nil
}

func (c *cardTransaction) CreatePairs(ctx context.Context, ids [2]int64) error {
	_, err := c.tx.ExecContext(
		ctx,
		fmt.Sprintf("INSERT INTO %s (origin_id, pair_with) VALUES ($1, $2), ($2, $1)", pairsTableName),
		ids[0],
		ids[1],
	)
	if err != nil {
		c.tx.Rollback()
	}
	return err
}

func (c *cardTransaction) GetPairsIDs(ctx context.Context, id int64) ([]int64, error) {
	query := fmt.Sprintf("SELECT pair_with FROM %s WHERE origin_id = %d", pairsTableName, id)
	var ids []int64
	err := c.tx.SelectContext(ctx, &ids, query)
	if err != nil {
		c.tx.Rollback()
	}
	return ids, err
}

func (c *cardTransaction) GetCard(ctx context.Context, id int64) (dto.Card, error) {
	query := fmt.Sprintf("SELECT id, value, description, lang, image_hash, added_at FROM %s WHERE id=%d", cardTableName, id)
	var card dto.Card
	err := c.tx.QueryRowxContext(ctx, query).
		Scan(&card.ID, &card.Value, &card.Description, &card.Lang, &card.ImageHash, &card.AddedAt)
	if err != nil {
		c.tx.Rollback()
		return card, err
	}
	return card, err
}

func (c *cardTransaction) GetIDs(ctx context.Context, lang dto.Language, limit int64) ([]int64, error) {
	var query string
	if limit > 0 {
		query = fmt.Sprintf("SELECT id FROM %s WHERE lang=$1 LIMIT %d", cardTableName, limit)
	} else {
		query = fmt.Sprintf("SELECT id FROM %s WHERE lang=$1", cardTableName)
	}
	res, err := c.tx.QueryContext(ctx, query, lang)
	if err != nil {
		c.tx.Rollback()
		return nil, err
	}
	defer res.Close()
	ids := make([]int64, 0, 2)
	for res.Next() {
		var i int64
		if err := res.Scan(&i); err == nil {
			ids = append(ids, i)
		}
	}
	return ids, nil
}

func (c *cardTransaction) End() error {
	return c.tx.Commit()
}
