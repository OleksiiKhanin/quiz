package api

import (
	"english-card/interfaces"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type Router struct {
	*mux.Router
}

func getLoggerMiddleware(log io.Writer) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func(start time.Time) {
				fmt.Fprintf(log, "time='%s' path='%s' method='%s' remote_addr='%s' duration='%s'",
					start.Format(time.RFC3339),
					r.URL.Path,
					r.Method,
					r.RemoteAddr,
					time.Since(start).String(),
				)
			}(time.Now())
			next.ServeHTTP(w, r)
		})
	}
}

func addJsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func getRecoverThePanicMiddleware(panicHandler func(any)) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				msg := recover()
				if msg != nil {
					panicHandler(msg)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func GetRouter(card interfaces.CardService, images interfaces.ImageService) *Router {
	r := mux.NewRouter()

	apiCard := GetCardAPI(card)
	apiImage := GetImageAPI(images)

	r.Use(addJsonContentTypeMiddleware)
	r.Use(getRecoverThePanicMiddleware(func(msg any) { fmt.Fprintf(os.Stderr, "PANIC covered by middleware: %s", msg) }))
	r.Use(getLoggerMiddleware(os.Stdout))

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
