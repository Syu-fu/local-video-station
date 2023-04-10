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

func SelectVideoCount(db *sql.DB) (int, error) {
	const sqlStr = `
		select count(1)
		from video
		;
	`

	row := db.QueryRow(sqlStr)
	if err := row.Err(); err != nil {
		return 0, row.Err()
	}

	var rowCount int
	if err := row.Scan(&rowCount); err != nil {
		return 0, err
	}

	return rowCount, nil
}

func SelectVideoDetail(db *sql.DB, id string) (models.Video, error) {
	const sqlStr = `
		select id, title, title_reading, url, thumbnail_url
		from video
		where id = ?
		;
	`

	row := db.QueryRow(sqlStr, id)
	if err := row.Err(); err != nil {
		return models.Video{}, err
	}

	var video models.Video

	if err := row.Scan(&video.ID, &video.Title, &video.TitleReading, &video.Url, &video.ThumbnailUrl); err != nil {
		return models.Video{}, err
	}

	return video, nil
}
