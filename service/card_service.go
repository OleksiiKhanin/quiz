package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"english-card/dto"
	"english-card/interfaces"
)

type cardService struct {
	repo   interfaces.CardRepo
	images interfaces.ImageRepo
}

func CreateCardService(repo interfaces.CardRepo, images interfaces.ImageRepo) interfaces.CardService {
	return &cardService{repo: repo, images: images}
}

func (c *cardService) AddCards(ctx context.Context, image *dto.Image, pairs ...[2]*dto.Card) error {
	cardTransaction, err := c.repo.Begin(nil)
	if err != nil {
		return fmt.Errorf("create a card transaction failed: %w", err)
	}
	defer cardTransaction.End()
	var hash string
	if image != nil {
		imageTransaction, err := c.images.Begin(cardTransaction)
		if err != nil {
			return err
		}
		if hash, err = imageTransaction.InsertImage(ctx, image.Title, image.Type, image.Data); err != nil {
			return err
		}
	}
	for _, cards := range pairs {
		for i := range cards {
			if cards[i].Value == "" {
				return errors.New("value parameter is required")
			}
			if hash != "" {
				cards[i].ImageHash = hash
			}
			cards[i].ID, err = cardTransaction.InsertCard(ctx, cards[i])
			if err != nil {
				return err
			}
		}
		err := cardTransaction.CreatePairs(ctx, [2]int64{cards[0].ID, cards[1].ID})
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *cardService) UpdateCard(ctx context.Context, card *dto.Card) error {
	cardTransaction, err := c.repo.Begin(nil)
	if err != nil {
		return fmt.Errorf("update a card transaction failed: %w", err)
	}
	defer cardTransaction.End()
	if err := cardTransaction.UpdateCard(ctx, card); err != nil {
		return fmt.Errorf("update card service failed with: %w", err)
	}
	return nil
}

func (c *cardService) GetCards(ctx context.Context, id int64) ([]dto.Card, error) {
	cardTransaction, err := c.repo.Begin(nil)
	if err != nil {
		return nil, fmt.Errorf("get card transaction failed: %w", err)
	}
	defer cardTransaction.End()
	origin, err := cardTransaction.GetCard(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get card ids failed:%w", err)
	}
	results := make([]dto.Card, 1)
	results[0] = origin
	ids, err := cardTransaction.GetPairsIDs(ctx, origin.ID)
	if err != nil {
		return results, fmt.Errorf("get compliance ids failed:%w", err)
	}
	for i := range ids {
		if card, err := cardTransaction.GetCard(ctx, ids[i]); err == nil {
			results = append(results, card)
		}
	}
	return results, nil
}

func (c *cardService) GetRandomCards(ctx context.Context, lang dto.Language) ([]dto.Card, error) {
	cardTransaction, err := c.repo.Begin(nil)
	if err != nil {
		return nil, err
	}
	defer cardTransaction.End()
	ids, err := cardTransaction.GetIDs(ctx, lang, 0)
	if err != nil {
		return nil, fmt.Errorf("get all ids specified language %s: %w", lang, err)
	}
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(ids))
	return c.GetCards(ctx, ids[i])
}
