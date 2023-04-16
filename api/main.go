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
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	dbUser        = os.Getenv("DB_USER")
	dbPassword    = os.Getenv("DB_PASSWORD")
	apiPort       = os.Getenv("API_PORT")
	dbDatabase    = "local_video_station"
	timeoutSecond = 30
	dbConn        = fmt.Sprintf("%s:%s@tcp(db)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	endpoint      = "minio:" + os.Getenv("MINIO_PORT")
	useSSL        = false
)

func main() {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println(dbConn)
		log.Println("fail to connect DB")

		return
	}

	mc, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4("", "", ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
		log.Println("fail to connect MinIO")

		return
	}

	r := api.NewRouter(db, mc)

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
