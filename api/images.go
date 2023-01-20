package api

import (
	"encoding/json"
	"english-card/dto"
	"english-card/interfaces"
	"net/http"

	"github.com/gorilla/mux"
)

type imageAPI struct {
	svc interfaces.ImageService
}

func GetImageAPI(svc interfaces.ImageService) imageAPI {
	return imageAPI{svc}
}

func (i *imageAPI) CreateImageHandler(w http.ResponseWriter, r *http.Request) {
	var image dto.Image
	err := json.NewDecoder(r.Body).Decode(&image)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(apiError{ErrMsg: err.Error()})
		return
	}
	img, err := i.svc.SaveImage(r.Context(), image.Title, image.Type, image.Data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(apiError{ErrMsg: err.Error()})
	} else {
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(img)
	}
}

func (i *imageAPI) GetImageDataHandler(w http.ResponseWriter, r *http.Request) {
	hash := mux.Vars(r)["hash"]
	image, err := i.svc.GetImage(r.Context(), hash)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(apiError{ErrMsg: err.Error()})
	} else {
		w.Header().Add("Content-Type", image.Type)
		w.WriteHeader(http.StatusOK)
		w.Write(image.Data)
	}
}

func (i *imageAPI) GetImageObjectHandler(w http.ResponseWriter, r *http.Request) {
	hash := mux.Vars(r)["hash"]
	image, err := i.svc.GetImage(r.Context(), hash)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(apiError{ErrMsg: err.Error()})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(image)
	}
}
