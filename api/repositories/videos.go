package repositories

import (
	"database/sql"
	"os"
	"strings"

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

		video.Url = "http://" + os.Getenv("IPADDRESS") + ":" + os.Getenv("MINIO_PORT") + "/data/" + video.Url
		video.ThumbnailUrl = "http://" + os.Getenv("IPADDRESS") + ":" + os.Getenv("MINIO_PORT") + "/data/" + video.ThumbnailUrl

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

	video.Url = "http://" + os.Getenv("IPADDRESS") + ":" + os.Getenv("MINIO_PORT") + "/data/" + video.Url
	video.ThumbnailUrl = "http://" + os.Getenv("IPADDRESS") + ":" + os.Getenv("MINIO_PORT") + "/data/" + video.ThumbnailUrl

	return video, nil
}

func InsertVideo(db *sql.DB, video models.Video) (models.Video, error) {
	const sqlStr = `
	insert into video (id, title, title_reading, url, thumbnail_url) values
	(?, ?, ?, ?, ?)
	;
	`

	_, err := db.Exec(sqlStr, video.ID, video.Title, video.TitleReading, video.Url, video.ThumbnailUrl)
	if err != nil {
		return models.Video{}, err
	}

	return video, nil
}

func SelectVideoListByTags(db *sql.DB, tagIDs string, page int) ([]models.Video, error) {
	tags := strings.Split(tagIDs, ",")
	numTags := len(tags)

	sqlStr := `
		SELECT v.id, v.title, v.title_reading, v.url, v.thumbnail_url
		FROM video v
		JOIN video_tag vt ON v.id = vt.video_id
		WHERE vt.tag_id IN (?)
		GROUP BY v.id
		HAVING COUNT(DISTINCT vt.tag_id) = ?
		ORDER BY v.updated_at
		LIMIT ? OFFSET ?;
	`

	offset := (page - 1) * videoNumPerPage

	rows, err := db.Query(sqlStr, tagIDs, numTags, videoNumPerPage, offset)
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

		video.Url = "http://" + os.Getenv("IPADDRESS") + ":" + os.Getenv("MINIO_PORT") + "/data/" + video.Url
		video.ThumbnailUrl = "http://" + os.Getenv("IPADDRESS") + ":" + os.Getenv("MINIO_PORT") + "/data/" + video.ThumbnailUrl

		videoArray = append(videoArray, video)
	}

	return videoArray, nil
}

func SelectVideoCountByTags(db *sql.DB, tagIDs string) (int, error) {
	tags := strings.Split(tagIDs, ",")
	numTags := len(tags)

	sqlStr := `
		select count(1)
		from video v
		join video_tag vt on v.id = vt.video_id
		where vt.tag_id in (?)
		group by v.id
		having count(distinct vt.tag_id) = ?
	`

	row := db.QueryRow(sqlStr, tagIDs, numTags)
	if err := row.Err(); err != nil {
		return 0, row.Err()
	}

	var rowCount int
	if err := row.Scan(&rowCount); err != nil {
		return 0, err
	}

	return rowCount, nil
}
