package services_test

import (
	"testing"

	"api/models"
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

func TestGetVideoCountService(t *testing.T) {
	expectedNum := len(testdata.VideoTestData)

	got, err := aSer.GetVideoCountService()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if num := got; num != expectedNum {
		t.Errorf("want %d but got %d videos\n", expectedNum, num)
	}
}

func TestGetVideoService(t *testing.T) {
	got, err := aSer.GetVideoService(testdata.VideoTestData[0].ID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !equalVideo(testdata.VideoTestData[0], got) {
		t.Errorf("want %v(%T) but got %v(%T)\n", testdata.VideoTestData[0], testdata.VideoTestData[0], got, got)
	}
}

func equalVideo(v1, v2 models.Video) bool {
	if v1.ID != v2.ID {
		return false
	}

	if v1.Title != v2.Title {
		return false
	}

	if v1.TitleReading != v2.TitleReading {
		return false
	}

	if v1.Url != v2.Url {
		return false
	}

	if v1.ThumbnailUrl != v2.ThumbnailUrl {
		return false
	}

	if len(v1.Tags) != len(v2.Tags) {
		return false
	}

	tagMap1 := make(map[models.Tag]bool)
	for _, tag := range v1.Tags {
		tagMap1[tag] = true
	}

	for _, tag := range v2.Tags {
		if _, ok := tagMap1[tag]; !ok {
			return false
		}
	}

	return true
}
