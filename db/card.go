package db

import (
	"context"
	"english-card/dto"
	"english-card/interfaces"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const cardTableName = "languages"
const complienceTableName = "compliances"

func GetCardRepo(db *sqlx.DB, images interfaces.ImageRepo) interfaces.CardRepo {
	return &cardRepo{db: db, images: images}
}

type cardRepo struct {
	db     *sqlx.DB
	images interfaces.ImageRepo
}

func (c *cardRepo) Begin() (interfaces.CardTx, error) {
	tx, err := c.db.Beginx()
	if err != nil {
		return nil, err
	}
	images, err := c.images.Begin()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return &cardTransaction{tx: tx, images: images}, nil
}

type cardTransaction struct {
	tx     *sqlx.Tx
	images interfaces.ImageTx
}

func (c *cardTransaction) InsertCard(ctx context.Context, card *dto.Card) (int64, error) {
	var err error
	if len(card.Image.Data) != 0 && card.Image.Hash == "" {
		card.Image.Hash, err = c.images.SaveImage(ctx, card.Image.Tittle, card.Image.Data)
		if err != nil {
			c.tx.Rollback()
			return 0, err
		}
	}
	result, err := c.tx.ExecContext(
		ctx,
		fmt.Sprintf("INSERT INTO %s (value, description, lang, image_uuid) VALUES ($1,$2,$3,$4)",
			cardTableName),
		card.Value,
		card.Description,
		card.Lang,
		card.Image.Hash,
	)
	if err != nil {
		c.tx.Rollback()
		return 0, err
	}
	return result.LastInsertId()
}

func (c *cardTransaction) CreateComplience(ctx context.Context, ids [2]int64) error {
	_, err := c.tx.ExecContext(
		ctx,
		fmt.Sprintf("INSERT INTO %s (origin_id, compliance_with) VALUES ($1, $2), ($2, $1)",
			complienceTableName,
		),
		ids[0],
		ids[1],
	)
	if err != nil {
		c.tx.Rollback()
	}
	return err
}

func (c *cardTransaction) GetComplianceCards(ctx context.Context, id int64) ([]dto.Card, error) {
	panic("implement me")
}

func (c *cardTransaction) GetCard(ctx context.Context, id int64) (dto.Card, error) {
	panic("implement me")
}
func (c *cardTransaction) GetIDs(ctx context.Context, lang dto.Language, limit int64) ([]int64, error) {
	res, err := c.tx.QueryContext(ctx,
		fmt.Sprintf("SELECT id FROM %s WHERE lang=$1 LIMIT $2", cardTableName),
		lang,
		limit,
	)
	if err != nil {
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
	err := c.images.End()
	if err != nil {
		c.tx.Rollback()
		return err
	}
	return c.tx.Commit()
}
