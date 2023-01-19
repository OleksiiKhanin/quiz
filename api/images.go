package api

import (
	"encoding/json"
	"english-card/dto"
	"english-card/interfaces"
	"net/http"
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
	img, err := i.svc.SaveImage(r.Context(), image.Tittle, image.Data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(apiError{ErrMsg: err.Error()})
	} else {
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(img)
	}
}
