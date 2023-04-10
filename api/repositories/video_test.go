package repositories_test

import (
	"reflect"
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

func TestSelectVideoCount(t *testing.T) {
	expectedNum := len(testdata.VideoTestData1)

	got, err := repositories.SelectVideoCount(testDB)
	if err != nil {
		t.Fatal(err)
	}

	if num := got; num != expectedNum {
		t.Errorf("want %d but got %d videos\n", expectedNum, num)
	}
}

func TestSelectVideoDetail(t *testing.T) {
	got, err := repositories.SelectVideoDetail(testDB, testdata.VideoTestData1[0].ID)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(testdata.VideoTestData1[0], got) {
		t.Errorf("want %v(%T) but got %v(%T) videos\n", testdata.VideoTestData1[0], testdata.VideoTestData1[0], got, got)
	}
}
