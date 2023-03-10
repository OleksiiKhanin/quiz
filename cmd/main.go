package main

import (
	"context"
	"english-card/api"
	"english-card/db"
	"english-card/service"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	storage := db.GetDB(opt.DB.Host, opt.DB.User, opt.DB.Pass, opt.DB.Name, opt.DB.Port)

	if err := db.MigrateSchema(storage.DB, opt.MigrationPath); err != nil {
		log.Fatalln(err.Error())
	}

	images := db.GetImagesRepo(storage)
	cards := db.GetCardRepo(storage)
	imageSVC := service.CreateImageService(images)
	cardSVC := service.CreateCardService(cards, images)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM)
	router := api.GetRouter(cardSVC, imageSVC)
	router.AddStatic("/static", opt.StaticFilesPath)
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", opt.Port),
		Handler: router,
	}

	go srv.ListenAndServe()
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Fatalln(srv.Shutdown(ctx))
}

var opt Options

type Options struct {
	Port            int
	MigrationPath   string
	StaticFilesPath string
	DB              struct {
		User string
		Pass string
		Host string
		Name string
		Port int
	}
}

func init() {
	flag.IntVar(&opt.Port, "p", 8080, "Sets the listening port")
	flag.IntVar(&opt.DB.Port, "dport", 5432, "Port where postgress database service is works")
	flag.StringVar(&opt.DB.Host, "dhost", "postgres", "Host where postgress database service is works")
	flag.StringVar(&opt.DB.Name, "dbase", "card_db", "Default database name")
	flag.StringVar(&opt.DB.User, "duser", "postgres", "Database user")
	flag.StringVar(&opt.DB.Pass, "dpass", "postgres", "Database password")
	flag.StringVar(&opt.MigrationPath, "m", "/migrations/", "Path to the migration files folder")
	flag.StringVar(&opt.StaticFilesPath, "s", "/static/", "Path to the static html, js, css files and icons, images etc")

	flag.Parse()
}
