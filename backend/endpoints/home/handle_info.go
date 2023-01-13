package home

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"backend/functions/profile"
	"backend/types/payload"
	"backend/types/response"
	"backend/utils/value"
)

func InfoHandler(c *fiber.Ctx) error {
	return c.JSON(response.Success(map[string]any{
		"overview": map[string]any{
			"name": "John Doe",
			"last_visited_topic": map[string]string{
				"title": "Responsive Design with Bootstrap",
				"href":  "https://gdsc.sit.kmutt.ac.th/bookstack/books/go-api-internet-kubernetes",
			},
		},
		"badge": map[string]int{
			"mention": 12,
			"quest":   16,
		},
		"monthly_progress": profile.GetProgressGraph(),
		"notifications": []*payload.HomeNotification{
			{
				Type: value.Ptr("flag"),
				Href: value.Ptr("https://www.google.com/"),
				Date: value.Ptr(time.Now().Add(-2 * time.Minute)),
				Messages: []*payload.HomeMessage{
					{
						Type: value.Ptr("link"),
						Href: value.Ptr("https://gdsc.community.dev/u/mrjt9j/"),
						Text: value.Ptr("B"),
					},
					{
						Type: value.Ptr("text"),
						Text: value.Ptr(" marked your comment as helpful "),
					},
				},
			},
			{
				Type: value.Ptr("comment"),
				Href: value.Ptr("https://www.google.com/"),
				Date: value.Ptr(time.Now().Add(-2 * 24 * time.Hour)),
				Messages: []*payload.HomeMessage{
					{
						Type: value.Ptr("link"),
						Href: value.Ptr("https://gdsc.community.dev/u/m43afg/"),
						Text: value.Ptr("E"),
					},
					{
						Type: value.Ptr("text"),
						Text: value.Ptr(" replied to your comment in "),
					},
					{
						Type: value.Ptr("link"),
						Href: value.Ptr("https://gdsc.community.dev/u/mwktp5/"),
						Text: value.Ptr("N"),
					},
					{
						Type: value.Ptr("text"),
						Text: value.Ptr("'s post "),
					},
				},
			},
		},
	}))
}
