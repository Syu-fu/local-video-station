package services

import (
	"database/sql"

	"github.com/minio/minio-go/v7"
)

type MyAppService struct {
	db *sql.DB
	mc *minio.Client
}

func NewMyAppService(db *sql.DB, mc *minio.Client) *MyAppService {
	return &MyAppService{db: db, mc: mc}
}
