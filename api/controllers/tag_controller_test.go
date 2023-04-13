package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestPostTagHandler(t *testing.T) {
	e := echo.New()
	tests := []struct {
		name       string
		tag        string
		resultCode int
	}{
		{name: "valid", tag: `{ "id": "", "name": "Sport", "nameReading": "Sport" }`, resultCode: http.StatusOK},
		{name: "invalid", tag: `{ id: "", }`, resultCode: http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := "http://localhost:8080/tag"
			req := httptest.NewRequest(http.MethodPost, url, strings.NewReader(tt.tag))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			res := httptest.NewRecorder()
			c := e.NewContext(req, res)
			err := tCon.PostTagHandler(c)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if res.Code != tt.resultCode {
				t.Errorf("unexpected StatusCode: want %d but %d\n", tt.resultCode, res.Code)
			}
		})
	}
}
