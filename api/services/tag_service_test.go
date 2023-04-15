package services_test

import (
	"testing"

	"api/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func TestPostTagService(t *testing.T) {
	tag := models.Tag{
		ID:          uuid.New().String(),
		Name:        "Sport",
		NameReading: "Sport(Reading)",
	}

	newTag, err := aSer.PostTagService(tag)
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
