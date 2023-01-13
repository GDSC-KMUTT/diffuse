package profile

import "backend/types/payload"

func GetProgressGraph() []*payload.ProfileProgressGraphItem {
	return []*payload.ProfileProgressGraphItem{
		{
			Name: "Thu",
			Val:  10,
		},
		{
			Name: "Fri",
			Val:  20,
		},
		{
			Name: "Sat",
			Val:  2,
		},
		{
			Name: "Sun",
			Val:  5,
		},
		{
			Name: "Mon",
			Val:  12,
		},
		{
			Name: "Tue",
			Val:  10,
		},
		{
			Name: "Wed",
			Val:  14,
		},
	}
}
