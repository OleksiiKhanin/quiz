package db

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"

	"english-card/dto"
	"english-card/interfaces"

	"github.com/jmoiron/sqlx"
)

const imageTableName = "images"

func GetImagesRepo(db *sqlx.DB) interfaces.ImageRepo {
	return &imagesRepo{db: db}
}

type imagesRepo struct {
	db *sqlx.DB
}

func (i *imagesRepo) Begin(t interfaces.TransactionGetter) (interfaces.ImageTx, error) {
	if t != nil {
		if tx := t.GetTX(); tx != nil {
			return &imagesTx{tx: tx}, nil
		}
	}
	tx, err := i.db.Beginx()
	if err != nil {
		return nil, err
	}
	return &imagesTx{tx: tx}, nil
}

type imagesTx struct {
	tx *sqlx.Tx
}

func (i *imagesTx) GetTX() *sqlx.Tx {
	return i.tx
}

func (i *imagesTx) GetImage(ctx context.Context, hash string) (dto.Image, error) {
	query := fmt.Sprintf("SELECT hash, type, title, data, added_at FROM %s WHERE hash=$1", imageTableName)
	var img = dto.Image{}
	err := i.tx.GetContext(ctx, &img, query, hash)
	return img, err
}

func (i *imagesTx) InsertImage(ctx context.Context, title, typ string, image []byte) (string, error) {
	hash := sha256.Sum256(image)
	uuid := strings.ReplaceAll(base64.StdEncoding.EncodeToString(hash[:]), "/", "_")
	query := fmt.Sprintf("INSERT INTO %s (hash, data, title, type) VALUES ($1,$2, $3, $4) ON CONFLICT DO NOTHING", imageTableName)
	if _, err := i.tx.ExecContext(ctx, query, uuid, image, title, typ); err != nil {
		i.tx.Rollback()
		return "", err
	}
	return uuid, nil
}

func (i *imagesTx) End() error {
	return i.tx.Commit()
}
