package service

import (
	"context"
	"english-card/dto"
	"english-card/interfaces"
)

type imageService struct {
	repo interfaces.ImageRepo
}

func CreateImageService(repo interfaces.ImageRepo) interfaces.ImageService {
	return &imageService{repo: repo}
}

func (i *imageService) GetImage(ctx context.Context, hash string) (dto.Image, error) {
	imageTx, err := i.repo.Begin(nil)
	if err != nil {
		return dto.Image{}, err
	}
	defer imageTx.End()
	return imageTx.GetImage(ctx, hash)
}

func (i *imageService) SaveImage(ctx context.Context, tittle, typ string, image []byte) (string, error) {
	imageTx, err := i.repo.Begin(nil)
	if err != nil {
		return "", err
	}
	defer imageTx.End()
	hash, err := imageTx.InsertImage(ctx, tittle, typ, image)
	if err != nil {
		return "", err
	}
	return hash, nil
}
