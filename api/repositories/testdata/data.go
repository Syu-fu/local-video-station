package testdata

import "api/models"

var VideoTestData1 = []models.Video{
	{
		ID:           "3C6FC606-3DC8-4EE5-AFC0-D6C20DE47AE4",
		Title:        "How to Make Sushi",
		TitleReading: "How to Make Sushi",
		Url:          "http://localhost:9000/data/video/how-to-make-sushi",
		ThumbnailUrl: "http://localhost:9000/data/thumbnail/sushi-thumbnail.jpg",
	},
	{
		ID:           "042243E2-9C87-4667-854A-AA1D6C8F64B8",
		Title:        "Hiking in the Mountains",
		TitleReading: "Hiking in the Mountains",
		Url:          "http://localhost:9000/data/video/hiking-in-the-mountains",
		ThumbnailUrl: "http://localhost:9000/data/thumbnail/hiking-thumbnail.jpg",
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

var TagTestData2 = []models.Tag{
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
	{
		ID:          "82AC868C-F677-4D08-B17C-F93B8913B69A",
		Name:        "Hiking",
		NameReading: "Hiking",
	},
	{
		ID:          "77BABEC0-6941-4F27-9094-56F400B90B1F",
		Name:        "Nature",
		NameReading: "Nature",
	},
	{
		ID:          "0431297E-0AB2-4E82-BE91-BB11AAB80C8B",
		Name:        "Mountains",
		NameReading: "Mountains",
	},
}
