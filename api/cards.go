package api

import (
	"encoding/json"
	"english-card/dto"
	"english-card/interfaces"
	"net/http"
)

type cardAPI struct {
	svc interfaces.CardService
}

func GetCardAPI(svc interfaces.CardService) cardAPI {
	return cardAPI{svc}
}

func (c *cardAPI) CreateCardPairHandler(w http.ResponseWriter, r *http.Request) {
	var cardPair [2]*dto.Card
	err := json.NewDecoder(r.Body).Decode(&cardPair)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(apiError{ErrMsg: err.Error()})
		return
	}
	err = c.svc.AddCards(r.Context(), cardPair)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(apiError{ErrMsg: err.Error()})
	} else {
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(cardPair)
	}
}
