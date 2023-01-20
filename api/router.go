package api

import (
	"english-card/interfaces"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	*mux.Router
}

func addJsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func GetRouter(card interfaces.CardService, images interfaces.ImageService) *Router {
	r := mux.NewRouter()

	apiCard := GetCardAPI(card)
	apiImage := GetImageAPI(images)

	r.Use(addJsonContentTypeMiddleware)

	r.Methods(http.MethodPost).Path("/v1/card").HandlerFunc(apiCard.CreateCardPairHandler)
	r.Methods(http.MethodPost).Path("/v1/image").HandlerFunc(apiImage.CreateImageHandler)

	r.Methods(http.MethodGet).Path("/v1/card/id/{id}").HandlerFunc(apiCard.GetCardsHandler)
	r.Methods(http.MethodGet).Path("/v1/card/lang/{lang}").HandlerFunc(apiCard.GetRandomCardsHandler)
	r.Methods(http.MethodGet).Path("/v1/image/data/{hash}").HandlerFunc(apiImage.GetImageDataHandler)
	r.Methods(http.MethodGet).Path("/v1/image/object/{hash}").HandlerFunc(apiImage.GetImageObjectHandler)

	return &Router{r}
}

func (r *Router) AddStatic(url, path string) {
	fileServer := http.FileServer(http.Dir(path))
	r.PathPrefix(url).Handler(http.StripPrefix(url, fileServer))
	r.Handle("/", http.RedirectHandler(url, http.StatusPermanentRedirect))
}
