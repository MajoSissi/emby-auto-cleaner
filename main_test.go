package main

import (
	"testing"
	"time"
)

func TestShouldDelete(t *testing.T) {
	config := &Config{}
	config.Cleanup.ProtectTags = []string{"keep"}
	config.Cleanup.ProtectFavorites = true

	watchedCutoff := time.Now().AddDate(0, 0, -30)

	tests := []struct {
		name     string
		item     EmbyItem
		expected bool
	}{
		{
			name: "已观看超过30天的剧集",
			item: EmbyItem{
				Type: "Episode",
				Path: "/path/to/episode.mp4",
				UserData: UserData{
					LastPlayedDate: time.Now().AddDate(0, 0, -40).Format(time.RFC3339),
					Played:         true,
					IsFavorite:     false,
				},
			},
			expected: true,
		},
		{
			name: "收藏的剧集",
			item: EmbyItem{
				Type: "Episode",
				Path: "/path/to/episode.mp4",
				UserData: UserData{
					LastPlayedDate: time.Now().AddDate(0, 0, -40).Format(time.RFC3339),
					Played:         true,
					IsFavorite:     true,
				},
			},
			expected: false,
		},
		{
			name: "包含保护标签的剧集",
			item: EmbyItem{
				Type: "Episode",
				Path: "/path/to/episode.mp4",
				Tags: []string{"keep"},
				UserData: UserData{
					LastPlayedDate: time.Now().AddDate(0, 0, -40).Format(time.RFC3339),
					Played:         true,
					IsFavorite:     false,
				},
			},
			expected: false,
		},
		{
			name: "未观看的剧集",
			item: EmbyItem{
				Type: "Episode",
				Path: "/path/to/episode.mp4",
				UserData: UserData{
					Played: false,
				},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := shouldDelete(tt.item, config, watchedCutoff)
			if result != tt.expected {
				t.Errorf("shouldDelete() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
