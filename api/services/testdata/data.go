package testdata

import "api/models"

var VideoTestData = []models.Video{
	{
		ID:           "3C6FC606-3DC8-4EE5-AFC0-D6C20DE47AE4",
		Title:        "How to Make Sushi",
		TitleReading: "How to Make Sushi",
		Url:          "http://localhost:9000/data/video/how-to-make-sushi",
		ThumbnailUrl: "http://localhost:9000/data/thumbnail/sushi-thumbnail.jpg",
		Tags: []models.Tag{
			{
				ID:          "74FF4BBA-18A5-4ABC-9164-4987021D411F",
				Name:        "Sushi",
				NameReading: "Sushi",
			},
			{
				ID:          "A082A0B6-63DE-4E97-BA30-B29319863D4F",
				Name:        "Cooking",
				NameReading: "Cooking",
			},
		},
	},
	{
		ID:           "09876-54321-09876-54321",
		Title:        "Hiking in the Mountains",
		TitleReading: "Hiking in the Mountains",
		Url:          "http://localhost:9000/data/video/hiking-in-the-mountains",
		ThumbnailUrl: "http://localhost:9000/data/thumbnail/hiking-thumbnail.jpg",
	},
}
