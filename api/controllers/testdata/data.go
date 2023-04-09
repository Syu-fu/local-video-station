package testdata

import "api/models"

var VideoTestData = []models.Video{
	{
		ID:           "3C6FC606-3DC8-4EE5-AFC0-D6C20DE47AE4",
		Title:        "How to Make Sushi",
		TitleReading: "How to Make Sushi",
		Url:          "https://example.com/how-to-make-sushi",
		ThumbnailUrl: "https://example.com/sushi-thumbnail.jpg",
	},
	{
		ID:           "042243E2-9C87-4667-854A-AA1D6C8F64B8",
		Title:        "Hiking in the Mountains",
		TitleReading: "Hiking in the Mountains",
		Url:          "https://example.com/hiking-in-the-mountains",
		ThumbnailUrl: "https://example.com/hiking-thumbnail.jpg",
	},
}

var TagTestData = []models.Tag{
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
}
