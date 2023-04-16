package repositories_test

import (
	"testing"

	"api/models"
	"api/repositories"
)

func TestInsertVideoTag(t *testing.T) {
	video := models.Video{
		ID:           "a7488274-dbfd-11ed-afa1-0242ac120002",
		Title:        "Cooking in Nature",
		TitleReading: "Cooking in Nature(Reading)",
		Url:          "video/cooking-in-nature",
		ThumbnailUrl: "thumbnail/cooking-in-nature.png",
		Tags: []models.Tag{
			{
				ID:          "A082A0B6-63DE-4E97-BA30-B29319863D4F",
				Name:        "Cooking",
				NameReading: "Cooking",
			},
			{
				ID:          "77BABEC0-6941-4F27-9094-56F400B90B1F",
				Name:        "Nature",
				NameReading: "Nature",
			},
		},
	}

	_, err := repositories.InsertVideo(testDB, video)
	if err != nil {
		t.Error(err)
	}

	newVideo, err := repositories.InsertVideoTag(testDB, video)
	if err != nil {
		t.Error(err)
	}

	if newVideo.ID != video.ID {
		t.Errorf("new insert id is expected %s but got %s\n", video.ID, newVideo.ID)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from video_tag
			where video_id = ?
			;
		`
		_, _ = testDB.Exec(sqlStr, video.ID)
		const sqlStr2 = `
			delete from video
			where id = ?
			;
		`
		_, _ = testDB.Exec(sqlStr2, video.ID)
	})
}
