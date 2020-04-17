package api

import (
	"fmt"
	"github.com/nickmurr/golang-api/api/db"
	"github.com/nickmurr/golang-api/api/router"
	"github.com/nickmurr/golang-api/config"
	"log"
	"net/http"
)

func Init() {
	config.Load()
	db.Load()

}

func Run() {

	Init()
	config.Log.Printf("Running API on port :%d", config.PORT)

	newRouter := router.NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.PORT), newRouter))
}
