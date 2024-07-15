package discourse

import (
	"testing"
)

func TestGenerateDirectoryQueryString(t *testing.T) {
	tests := []struct {
		name  string
		query UserDirectoryQuery
		want  string
	}{
		{
			"All fields empty",
			UserDirectoryQuery{
				Period:    "",
				Order:     "",
				Ascending: false,
				Page:      0,
			},
			"period=all",
		},
		{
			"Defaults",
			UserDirectoryQuery{},
			"period=all",
		},
		{
			"All fields",
			UserDirectoryQuery{
				Period:    PeriodMonthly,
				Order:     OrderLatestTopic,
				Ascending: true,
				Page:      32,
			},
			"period=monthly&order=latest_topic&asc=true&page=32",
		},
		{
			"Ascend",
			UserDirectoryQuery{
				Ascending: true,
			},
			"period=all&asc=true",
		},
		{
			"Additional options",
			UserDirectoryQuery{
				Period: PeriodDaily,
				Order:  OrderLikesGiven,
			},
			"period=daily&order=likes_given",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateListUsersQuery(&tt.query)
			if result != tt.want {
				t.Errorf("received \"%s\", want \"%s\"", result, tt.want)
			}
		})
	}
}
