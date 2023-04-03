package testdata

import "api/models"

var VideoTestData1 = []models.Video{
	{
		ID:           "12345-67890-12345-67890",
		Title:        "How to Make Sushi",
		TitleReading: "How to Make Sushi",
		Url:          "https://example.com/how-to-make-sushi",
		ThumbnailUrl: "https://example.com/sushi-thumbnail.jpg",
	},
	{
		ID:           "09876-54321-09876-54321",
		Title:        "Hiking in the Mountains",
		TitleReading: "Hiking in the Mountains",
		Url:          "https://example.com/hiking-in-the-mountains",
		ThumbnailUrl: "https://example.com/hiking-thumbnail.jpg",
	},
}
