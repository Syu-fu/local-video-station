package controllers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestVideoListHandler(t *testing.T) {
	e := echo.New()
	tests := []struct {
		name       string
		query      string
		resultCode int
	}{
		{name: "number query", query: "1", resultCode: http.StatusOK},
		{name: "alphabet query", query: "aaa", resultCode: http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("http://localhost:8080/video/list?page=%s", tt.query)
			req := httptest.NewRequest(http.MethodGet, url, nil)

			res := httptest.NewRecorder()
			c := e.NewContext(req, res)
			err := aCon.VideoListHandler(c)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if res.Code != tt.resultCode {
				t.Errorf("unexpected StatusCode: want %d but %d\n", tt.resultCode, res.Code)
			}
		})
	}
}

func TestVideoCountHandler(t *testing.T) {
	e := echo.New()
	tests := []struct {
		name       string
		resultCode int
	}{
		{name: "number query", resultCode: http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "http://localhost:8080/video/count"
			req := httptest.NewRequest(http.MethodGet, url, nil)

			res := httptest.NewRecorder()
			c := e.NewContext(req, res)
			err := aCon.VideoCountHandler(c)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestVideoDetailHandler(t *testing.T) {
	e := echo.New()
	tests := []struct {
		name       string
		id         string
		resultCode int
	}{
		{name: "number query", id: "3C6FC606-3DC8-4EE5-AFC0-D6C20DE47AE4", resultCode: http.StatusOK},
		{name: "number query", id: "00000000-0000-0000-0000-000000000000", resultCode: http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("http://localhost:8080/video/%s", tt.id)
			req := httptest.NewRequest(http.MethodGet, url, nil)

			res := httptest.NewRecorder()
			c := e.NewContext(req, res)
			err := aCon.VideoDetailHandler(c)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
