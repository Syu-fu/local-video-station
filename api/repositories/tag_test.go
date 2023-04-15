package repositories_test

import (
	"reflect"
	"sort"
	"testing"

	"api/models"
	"api/repositories"
	"api/repositories/testdata"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
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

func TestInsertTag(t *testing.T) {
	tag := models.Tag{
		ID:          uuid.New().String(),
		Name:        "Sport",
		NameReading: "Sport(Reading)",
	}

	newTag, err := repositories.InsertTag(testDB, tag)
	if err != nil {
		t.Error(err)
	}

	if newTag.ID != tag.ID {
		t.Errorf("new tag id is expected %s but got %s\n", tag.ID, newTag.ID)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from tag
			where name = ? and name_reading = ?
		`
		_, _ = testDB.Exec(sqlStr, tag.Name, tag.NameReading)
	})
}
