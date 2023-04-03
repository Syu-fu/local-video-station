package repositories_test

import (
	"testing"

	"api/repositories"
	"api/repositories/testdata"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectVideoList(t *testing.T) {
	expectedNum := len(testdata.VideoTestData1)

	got, err := repositories.SelectVideoList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d videos\n", expectedNum, num)
	}
}
