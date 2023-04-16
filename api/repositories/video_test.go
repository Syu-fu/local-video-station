package repositories_test

import (
	"reflect"
	"testing"

	"api/models"
	"api/repositories"
	"api/repositories/testdata"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
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

func TestInsertVideo(t *testing.T) {
	video := models.Video{
		ID:           uuid.New().String(),
		Title:        "Play baseball",
		TitleReading: "Play baseball(Reading)",
		Url:          "video/play-baseball",
		ThumbnailUrl: "thumbnail/play-baseball.png",
	}

	newVideo, err := repositories.InsertVideo(testDB, video)
	if err != nil {
		t.Error(err)
	}

	if newVideo.ID != video.ID {
		t.Errorf("new video id is expected %s but got %s\n", video.ID, newVideo.ID)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from video
			where title = ? and title_reading = ? and url = ? and thumbnail = ?
			;
		`
		_, _ = testDB.Exec(sqlStr, video.Title, video.TitleReading, video.Url, video.ThumbnailUrl)
	})
}
