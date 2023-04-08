package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"api/api"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser        = os.Getenv("DB_USER")
	dbPassword    = os.Getenv("DB_PASSWORD")
	apiPort       = os.Getenv("API_PORT")
	dbDatabase    = "local_video_station"
	timeoutSecond = 30
	dbConn        = fmt.Sprintf("%s:%s@tcp(db)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func main() {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println(dbConn)
		log.Println("fail to connect DB")

		return
	}

	r := api.NewRouter(db)

	log.Printf("server start at port %s", apiPort)
	srv := &http.Server{
		Addr:              ":" + apiPort,
		Handler:           r,
		ReadHeaderTimeout: time.Duration(timeoutSecond) * time.Second,
	}

	if err = srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
