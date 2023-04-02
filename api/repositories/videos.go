package repositories

import (
	"database/sql"

	"api/models"
)

const (
	videoNumPerPage = 10
)

func SelectVideoList(db *sql.DB, page int) ([]models.Video, error) {
	const sqlStr = `
		select id, title, title_reading, url, thumbnail_url
		from video
		order by updated_at
		limit ? offset ?
		;
	`

	rows, err := db.Query(sqlStr, videoNumPerPage, ((page - 1) * videoNumPerPage))
	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, rows.Err()
	}

	defer rows.Close()

	videoArray := make([]models.Video, 0)

	for rows.Next() {
		var video models.Video
		if err := rows.Scan(&video.ID, &video.Title, &video.TitleReading, &video.Url, &video.ThumbnailUrl); err != nil {
			return nil, err
		}

		videoArray = append(videoArray, video)
	}

	return videoArray, nil
}
