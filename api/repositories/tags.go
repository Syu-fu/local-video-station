package repositories

import (
	"database/sql"

	"api/models"
)

func SelectTagList(db *sql.DB, videoID string) ([]models.Tag, error) {
	const sqlStr = `
		select id, name, name_reading
		from tag
		join video_tag
		where tag.id = video_tag.tag_id
		and video_tag.video_id = ?
		;
	`

	rows, err := db.Query(sqlStr, videoID)
	if err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	defer rows.Close()

	tagArray := make([]models.Tag, 0)

	for rows.Next() {
		var tag models.Tag
		if err := rows.Scan(&tag.ID, &tag.Name, &tag.NameReading); err != nil {
			return nil, err
		}

		tagArray = append(tagArray, tag)
	}

	return tagArray, nil
}

func InsertTag(db *sql.DB, tag models.Tag) (models.Tag, error) {
	const sqlStr = `
	insert into tag (id, name, name_reading) values
	(?, ?, ?)
	;
	`

	_, err := db.Exec(sqlStr, tag.ID, tag.Name, tag.NameReading)
	if err != nil {
		return models.Tag{}, err
	}

	return tag, nil
}

func SelectAllTags(db *sql.DB) ([]models.Tag, error) {
	const sqlStr = `
		select id, name, name_reading
		from tag
		;
	`

	rows, err := db.Query(sqlStr)
	if err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	defer rows.Close()

	tagArray := make([]models.Tag, 0)

	for rows.Next() {
		var tag models.Tag
		if err := rows.Scan(&tag.ID, &tag.Name, &tag.NameReading); err != nil {
			return nil, err
		}

		tagArray = append(tagArray, tag)
	}

	return tagArray, nil
}
