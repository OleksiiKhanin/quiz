package db

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

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

func (i *imagesRepo) Begin() (interfaces.ImageTx, error) {
	tx, err := i.db.Beginx()
	if err != nil {
		return nil, err
	}
	return &imagesTx{tx: tx}, nil
}

type imagesTx struct {
	tx *sqlx.Tx
}

func (i *imagesTx) GetImage(ctx context.Context, hash string) (dto.Image, error) {
	query := fmt.Sprintf("SELECT hash, tittle, data FROM %s WHERE hash=$1", imageTableName)
	var img = dto.Image{
		Hash: hash,
	}
	err := i.tx.QueryRowx(query, hash).StructScan(&img)
	return img, err
}

func (i *imagesTx) SaveImage(ctx context.Context, tittle string, image []byte) (string, error) {
	hash := sha256.Sum256(image)
	uuid := base64.StdEncoding.EncodeToString(hash[:])
	query := fmt.Sprintf("INSERT INTO %s (hash, data, tittle) VALUES ($1,$2, $3)", imageTableName)
	if _, err := i.tx.ExecContext(ctx, query, uuid, image, tittle); err != nil {
		i.tx.Rollback()
		return "", err
	}
	return uuid, nil
}

func (i *imagesTx) End() error {
	return i.tx.Commit()
}
