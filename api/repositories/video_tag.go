package repositories

import (
	"database/sql"

	"api/models"
)

func InsertVideoTag(db *sql.DB, video models.Video) (models.Video, error) {
	const sqlStr = `
	insert into video_tag (video_id, tag_id) values
	(?, ?)
	;
	`

	for _, tag := range video.Tags {
		_, err := db.Exec(sqlStr, video.ID, tag.ID)
		if err != nil {
			return models.Video{}, err
		}
	}

	return video, nil
}

func DeleteVideoTagsByVideoID(db *sql.DB, videoID string) (string, error) {
	sqlStr := `
	delete from video_tag where video_id=?
	;
	`

	_, err := db.Exec(sqlStr, videoID)
	if err != nil {
		return "", err
	}

	return videoID, err
}
