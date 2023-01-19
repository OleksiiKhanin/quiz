package service

import (
	"context"
	"english-card/dto"
	"english-card/interfaces"
	"fmt"
)

type cardService struct {
	repo interfaces.CardRepo
}

func CreateCardService(repo interfaces.CardRepo) interfaces.CardService {
	return &cardService{repo: repo}
}

func (c *cardService) AddCards(ctx context.Context, pairs ...[2]*dto.Card) error {
	cardTransaction, err := c.repo.Begin()
	if err != nil {
		return fmt.Errorf("create a card transaction failed: %w", err)
	}
	defer cardTransaction.End()
	for _, cards := range pairs {
		for i := range cards {
			cards[i].ID, err = cardTransaction.InsertCard(ctx, cards[i])
			if err != nil {
				return err
			}
		}
		err := cardTransaction.CreateComplience(ctx, [2]int64{cards[0].ID, cards[1].ID})
		if err != nil {
			return err
		}
	}
	return nil
}
func (c *cardService) GetCards(ctx context.Context, id int64) ([]dto.Card, error) {
	cardTransaction, err := c.repo.Begin()
	if err != nil {
		return nil, fmt.Errorf("get card transaction failed: %w", err)
	}
	defer cardTransaction.End()
	origin, err := cardTransaction.GetCard(ctx, id)
	if err != nil {
		return nil, err
	}
	results := make([]dto.Card, 1)
	results[0] = origin
	complience, err := cardTransaction.GetComplianceCards(ctx, origin.ID)
	if err != nil {
		return results, err
	}
	results = append(results, complience...)
	return results, nil
}

func (c *cardService) GetRandomCards(ctx context.Context, lang dto.Language) ([]dto.Card, error) {
	panic("implement me")
}
