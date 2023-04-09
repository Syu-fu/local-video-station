package repositories_test

import (
	"reflect"
	"sort"
	"testing"

	"api/repositories"
	"api/repositories/testdata"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectTagList(t *testing.T) {
	expectedNum := len(testdata.TagTestData)

	got, err := repositories.SelectTagList(testDB, "3C6FC606-3DC8-4EE5-AFC0-D6C20DE47AE4")
	if err != nil {
		t.Fatal(err)
	}

	sort.Slice(got, func(i, j int) bool {
		return got[i].ID > got[j].ID
	})

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d tags\n", expectedNum, num)
	}

	if reflect.DeepEqual(testdata.TagTestData, got) {
		t.Errorf("want %v (%T) but got %v(%T)\n", testdata.TagTestData, testdata.TagTestData, got, got)
	}
}
