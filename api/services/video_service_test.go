package services_test

import (
	"testing"

	"api/services/testdata"

	_ "github.com/go-sql-driver/mysql"
)

func TestGetVideoListService(t *testing.T) {
	expectedNum := len(testdata.VideoTestData)

	got, err := aSer.GetVideoListService(1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d videos\n", expectedNum, num)
	}
}
