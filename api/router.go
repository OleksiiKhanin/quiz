package api

import (
	"english-card/interfaces"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	*mux.Router
}

func GetRouter(card interfaces.CardService, images interfaces.ImageService) *Router {
	r := mux.NewRouter()

	apiCard := GetCardAPI(card)
	apiImage := GetImageAPI(images)

	r.Methods("POST").Path("/v1/card").HandlerFunc(apiCard.CreateCardPairHandler)
	r.Methods("POST").Path("/v1/image").HandlerFunc(apiImage.CreateImageHandler)

	return &Router{r}
}

func (r *Router) AddStatic(url, path string) {
	fileServer := http.FileServer(http.Dir(path))
	r.PathPrefix(url).Handler(http.StripPrefix(url, fileServer))
	r.Handle("/", http.RedirectHandler(url, http.StatusPermanentRedirect))
}
