package api

import (
	"encoding/json"
	"english-card/dto"
	"english-card/interfaces"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type cardAPI struct {
	svc interfaces.CardService
}

func GetCardAPI(svc interfaces.CardService) cardAPI {
	return cardAPI{svc}
}

func (c *cardAPI) CreateCardPairHandler(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateCardsRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(apiError{ErrMsg: err.Error()})
		return
	}
	err = c.svc.AddCards(r.Context(), &dto.Image{Data: request.ImageData, Title: request.ImageTittle, Type: request.ImageType}, request.Cards)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(apiError{ErrMsg: err.Error()})
	} else {
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(request.Cards)
	}
}

func (c *cardAPI) GetCardsHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(apiError{ErrMsg: err.Error()})
		return
	}
	cards, err := c.svc.GetCards(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(apiError{ErrMsg: err.Error()})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(cards)
	}
}

func (c *cardAPI) GetRandomCardsHandler(w http.ResponseWriter, r *http.Request) {
	lang := mux.Vars(r)["lang"]
	cards, err := c.svc.GetRandomCards(r.Context(), dto.Language(lang))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(apiError{ErrMsg: err.Error()})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(cards)
	}
}
